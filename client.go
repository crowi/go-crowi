// Package crowi provides some Crowi's APIs
package crowi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

const version = "0.1"

var userAgent = fmt.Sprintf("CrowiGoClient/%s (%s)", version, runtime.Version())

const (
	apiPagesCreate    = "/_api/pages.create"
	apiPagesUpdate    = "/_api/pages.update"
	apiAttachmentsAdd = "/_api/attachments.add"
)

type Client struct {
	URL        *url.URL
	Token      string
	HTTPClient *http.Client
}

// NewClient...
func NewClient(apiURL, token string) (*Client, error) {
	if len(apiURL) == 0 {
		return nil, errors.New("missing api url")
	}

	parsedURL, err := url.ParseRequestURI(apiURL)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse url: %s", apiURL)
	}

	if len(token) == 0 {
		return nil, errors.New("missing token")
	}

	return &Client{
		URL:        parsedURL,
		Token:      token,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}, nil
}

func (c *Client) newRequest(method, resource string, data url.Values) (*http.Request, error) {
	c.URL.Path = resource
	urlStr := fmt.Sprintf("%v", c.URL)

	req, err := http.NewRequest(method, urlStr, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Length", strconv.Itoa(len(data.Encode())))

	return req, nil
}

// CreatePage...
func (c *Client) CreatePage(path, body string) (*Pages, error) {
	data := url.Values{}
	data.Set("access_token", c.Token)
	data.Set("path", path)
	data.Set("body", body)

	req, err := c.newRequest("POST", apiPagesCreate, data)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var p Pages
	if err := decodeBody(res, &p); err != nil {
		return nil, err
	}

	return &p, nil
}

// UpdatePage...
func (c *Client) UpdatePage(pageID, body string) (*Pages, error) {
	data := url.Values{}
	data.Set("access_token", c.Token)
	data.Set("page_id", pageID)
	data.Set("body", body)

	req, err := c.newRequest("POST", apiPagesUpdate, data)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var p Pages
	if err := decodeBody(res, &p); err != nil {
		return nil, err
	}

	return &p, nil
}

func (c *Client) fileUpload(method, resource string, params map[string]string, filePath string) (*http.Request, error) {
	c.URL.Path = resource
	urlStr := fmt.Sprintf("%v", c.URL)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	w := multipart.NewWriter(body)
	part, err := w.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		err := w.WriteField(key, val)
		if err != nil {
			return nil, err
		}
	}

	err = w.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", w.FormDataContentType())
	req.Header.Set("User-Agent", userAgent)

	return req, nil
}

func (c *Client) fileUpload2(method, resource string, params map[string]string, filePath string) (*http.Request, error) {
	c.URL.Path = resource
	urlStr := fmt.Sprintf("%v", c.URL)

	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	for key, val := range params {
		err := writer.WriteField(key, val)
		if err != nil {
			return nil, err
		}
	}
	{
		header := make(textproto.MIMEHeader)
		header.Add("Content-Disposition", "form-data; name=\"file\"; filename=\"gopher.png\"")
		header.Add("Content-Type", "image/png")
		fileWriter, err := writer.CreatePart(header)
		if err != nil {
			return nil, err
		}
		file, err := os.Open("gopher.png")
		if err != nil {
			return nil, err
		}
		io.Copy(fileWriter, file)
	}
	writer.Close()

	req, err := http.NewRequest(method, urlStr, &buffer)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "multipart/form-data; boundary="+writer.Boundary())

	return req, nil
}

// AddAttachment...
func (c *Client) AddAttachment(pageID, filePath string) (*Attachments, error) {
	req, err := c.fileUpload2("POST", apiAttachmentsAdd, map[string]string{
		"access_token": c.Token,
		"page_id":      pageID,
	}, filePath)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var a Attachments
	if err := decodeBody(res, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}

// Package crowi provides some Crowi APIs for Go
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

type API interface {
	PagesCreate() (*Crowi, error)
	PagesUpdate() (*Crowi, error)
	AttachmentsAdd() (*Crowi, error)
}

// Client wraps http client
type Client struct {
	URL        *url.URL
	Token      string
	HTTPClient *http.Client
}

// NewClient creates an API client
func NewClient(apiURL, token string) (*Client, error) {
	if len(apiURL) == 0 {
		return nil, errors.New("missing api url")
	}

	if len(token) == 0 {
		return nil, errors.New("missing token")
	}

	parsedURL, err := url.ParseRequestURI(apiURL)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse url: %s", apiURL)
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

// PagesCreate makes a page in your Crowi. The request requires
// the path and page content used for the page name
func (c *Client) PagesCreate(path, body string) (*Crowi, error) {
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

	var crowi Crowi
	if err := decodeBody(res, &crowi); err != nil {
		return nil, err
	}

	return &crowi, nil
}

// PagesUpdate updates the page content. A page_id is necessary to know which
// page should be updated.
func (c *Client) PagesUpdate(pageID, body string) (*Crowi, error) {
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

	var crowi Crowi
	if err := decodeBody(res, &crowi); err != nil {
		return nil, err
	}

	return &crowi, nil
}

func (c *Client) fileUpload(method, resource string, params map[string]string, filePath string) (*http.Request, error) {
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
		header.Add("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="%s"`, filePath))
		header.Add("Content-Type", "image/png")
		fileWriter, err := writer.CreatePart(header)
		if err != nil {
			return nil, err
		}
		file, err := os.Open(filePath)
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
	req.Header.Set("User-Agent", userAgent)

	return req, nil
}

// AttachmentsAdd attaches an image file to the page. This request requires
// page_id and the image file path which you want to attach.
func (c *Client) AttachmentsAdd(pageID, filePath string) (*Crowi, error) {
	req, err := c.fileUpload("POST", apiAttachmentsAdd, map[string]string{
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

	var crowi Crowi
	if err := decodeBody(res, &crowi); err != nil {
		return nil, err
	}

	return &crowi, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}

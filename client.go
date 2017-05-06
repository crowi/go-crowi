package crowi

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
	"path"
	"runtime"
	"strings"
)

const version = "0.1"

var userAgent = fmt.Sprintf("CrowiGoClient/%s (%s)", version, runtime.Version())

type Client struct {
	http.Client

	config Config

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	Pages       *PagesService
	Attachments *AttachmentsService
}

type service struct {
	client *Client
}

type ListOptions struct {
	Pagenation bool
}

type Config struct {
	URL                string
	Token              string
	InsecureSkipVerify bool
}

func NewClient(cfg Config) (*Client, error) {
	if len(cfg.URL) == 0 {
		return nil, errors.New("missing api url")
	}

	if len(cfg.Token) == 0 {
		return nil, errors.New("missing token")
	}

	client := *http.DefaultClient
	if cfg.InsecureSkipVerify {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client = http.Client{Transport: tr}
	}

	c := &Client{
		Client: client,
		config: cfg,
	}
	c.common.client = c
	c.Pages = (*PagesService)(&c.common)
	c.Attachments = (*AttachmentsService)(&c.common)

	return c, nil
}

func (c *Client) newRequest(ctx context.Context, method string, uri string, params interface{}, res interface{}) error {
	u, err := url.Parse(c.config.URL)
	if err != nil {
		return err
	}
	u.Path = path.Join(u.Path, uri)

	values, ok := params.(url.Values)
	if !ok {
		return nil
	}

	var req *http.Request
	var body io.Reader
	if method == http.MethodGet {
		u.RawQuery = values.Encode()
	} else {
		body = strings.NewReader(values.Encode())
	}
	req, err = http.NewRequest(method, u.String(), body)
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", userAgent)
	if params != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return parseAPIError("bad request", resp)
	} else if res == nil {
		return nil
	}
	return json.NewDecoder(resp.Body).Decode(&res)
}

func (c *Client) newRequestWithFile(ctx context.Context, method string, uri string, params interface{}, res interface{}, file string) error {
	u, err := url.Parse(c.config.URL)
	if err != nil {
		return err
	}
	u.Path = path.Join(u.Path, uri)

	values, ok := params.(map[string]string)
	if !ok {
		return nil
	}

	var req *http.Request
	var buf bytes.Buffer
	mw := multipart.NewWriter(&buf)
	for key, val := range values {
		err := mw.WriteField(key, val)
		if err != nil {
			return err
		}
	}
	header := make(textproto.MIMEHeader)
	header.Add("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="%s"`, file))
	header.Add("Content-Type", "image/png")
	fileWriter, err := mw.CreatePart(header)
	if err != nil {
		return err
	}
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	io.Copy(fileWriter, f)
	mw.Close()

	req, err = http.NewRequest(method, u.String(), &buf)
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)
	req.Header.Add("Content-Type", "multipart/form-data; boundary="+mw.Boundary())
	req.Header.Set("User-Agent", userAgent)

	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return parseAPIError("bad request", resp)
	} else if res == nil {
		return nil
	}
	return json.NewDecoder(resp.Body).Decode(&res)
}

func parseAPIError(prefix string, resp *http.Response) error {
	errMsg := fmt.Sprintf("%s: %s", prefix, resp.Status)
	var e struct {
		Error string `json:"error"`
	}

	json.NewDecoder(resp.Body).Decode(&e)
	if e.Error != "" {
		errMsg = fmt.Sprintf("%s: %s", errMsg, e.Error)
	}

	return errors.New(errMsg)
}

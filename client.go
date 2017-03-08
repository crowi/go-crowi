// Package crowi provides some Crowi's APIs
package crowi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"runtime"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

const version = "0.1"

var userAgent = fmt.Sprintf("CrowiGoClient/%s (%s)", version, runtime.Version())

const (
	PagesCreateAPI    = "/_api/pages.create"
	PagesUpdateAPI    = "/_api/pages.update"
	AttachmentsAddAPI = "/_api/attachments.add"
)

type Client struct {
	Token      string
	URL        *url.URL
	HTTPClient *http.Client
}

// NewClient...
func NewClient(token, apiURL string) (*Client, error) {
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
		Token:      token,
		URL:        parsedURL,
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
func (c *Client) CreatePage(path, body string) (*CrowiPages, error) {
	data := url.Values{}
	data.Set("access_token", c.Token)
	data.Set("path", path)
	data.Set("body", body)

	req, err := c.newRequest("POST", PagesCreateAPI, data)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var crowi CrowiPages
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

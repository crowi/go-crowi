package crowi

import (
	"context"
	"net/http"
	"net/url"
	"time"
)

func (c *Client) PagesCreate(ctx context.Context, path, body string) (*Page, error) {
	var page Page
	params := url.Values{}
	params.Set("access_token", c.Token)
	params.Set("path", path)
	params.Set("body", body)
	err := c.newRequest(ctx, http.MethodPost, "/_api/pages.create", params, &page)
	if err != nil {
		return nil, err
	}
	return &page, nil
}

func (c *Client) PagesUpdate(ctx context.Context, id, body string) (*Page, error) {
	var page Page
	params := url.Values{}
	params.Set("access_token", c.Token)
	params.Set("page_id", id)
	params.Set("body", body)
	err := c.newRequest(ctx, http.MethodPost, "/_api/pages.update", params, &page)
	if err != nil {
		return nil, err
	}
	return &page, nil
}

func (c *Client) PagesList(ctx context.Context, path, user string) (*Pages, error) {
	var pages Pages
	params := url.Values{}
	params.Set("access_token", c.Token)
	params.Set("path", path)
	params.Set("user", user)
	err := c.newRequest(ctx, http.MethodGet, "/_api/pages.list", params, &pages)
	if err != nil {
		return nil, err
	}
	return &pages, nil
}

type Page struct {
	Page  PagesPage `json:"page"`
	OK    bool      `json:"ok"`
	Error string    `json:"error"`
}

type PagesPage struct {
	Revision       PageRevision  `json:"revision"`
	V              int           `json:"__v"`
	RedirectTo     interface{}   `json:"redirectTo"`
	UpdatedAt      time.Time     `json:"updatedAt"`
	LastUpdateUser interface{}   `json:"lastUpdateUser"`
	Creator        interface{}   `json:"creator"`
	Path           string        `json:"path"`
	ID             string        `json:"_id"`
	CreatedAt      time.Time     `json:"createdAt"`
	CommentCount   int           `json:"commentCount"`
	SeenUsers      []interface{} `json:"seenUsers"`
	Liker          []interface{} `json:"liker"`
	GrantedUsers   []string      `json:"grantedUsers"`
	Grant          int           `json:"grant"`
	Status         string        `json:"status"`
	Extended       PageExtended  `json:"extended,omitempty"`
}

type Pages struct {
	Pages []PagesPage `json:"pages"`
	OK    bool        `json:"ok"`
	Error string      `json:"error"`
}

type PageRevision struct {
	ID        string      `json:"_id"`
	Author    interface{} `json:"author"`
	Body      string      `json:"body"`
	Path      string      `json:"path"`
	V         int         `json:"__v"`
	CreatedAt time.Time   `json:"createdAt"`
	Format    string      `json:"format"`
}

type PageExtended struct {
	Slack string `json:"slack"`
}

package crowi

import (
	"context"
	"net/http"
	"net/url"
	"time"
)

func (c *Client) AttachmentsAdd(ctx context.Context, id, file string) (*Attachments, error) {
	var at Attachments
	params := map[string]string{
		"access_token": c.Token,
		"page_id":      id,
	}
	err := c.newRequestWithFile(ctx, http.MethodPost, "/_api/attachments.add", params, &at, file)
	if err != nil {
		return nil, err
	}
	return &at, nil
}

func (c *Client) AttachmentsList(ctx context.Context, id string) (*Attachments, error) {
	var at Attachments
	params := url.Values{}
	params.Set("access_token", c.Token)
	params.Set("page_id", id)
	err := c.newRequest(ctx, http.MethodGet, "/_api/attachments.list", params, &at)
	if err != nil {
		return nil, err
	}
	return &at, nil
}

type Attachments struct {
	Attachments []Attachment `json:"attachments"`
	OK          bool         `json:"ok"`
}

type Attachment struct {
	ID           string             `json:"_id"`
	FileFormat   string             `json:"fileFormat"`
	FileName     string             `json:"fileName"`
	OriginalName string             `json:"originalName"`
	FilePath     string             `json:"filePath"`
	Creator      AttachmentsCreator `json:"creator"`
	Page         string             `json:"page"`
	V            int                `json:"__v"`
	CreatedAt    time.Time          `json:"createdAt"`
	FileSize     int                `json:"fileSize"`
	URL          string             `json:"url"`
}

type AttachmentsCreator struct {
	ID        string    `json:"_id"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	V         int       `json:"__v"`
	APIToken  string    `json:"apiToken"`
	Admin     bool      `json:"admin"`
	CreatedAt time.Time `json:"createdAt"`
	Status    int       `json:"status"`
	Lang      string    `json:"lang"`
}

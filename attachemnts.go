package crowi

import (
	"context"
	"net/http"
	"net/url"
	"time"
)

// AttachmentsService handles communication with the Attachments related
// methods of the Crowi API.
type AttachmentsService service

// Add attaches an image file to the page. This request requires
// page_id and the image file path which you want to attach.
func (s *AttachmentsService) Add(ctx context.Context, id, file string) (*Attachment, error) {
	var at Attachment
	params := map[string]string{
		"access_token": s.client.Token,
		"page_id":      id,
	}
	err := s.client.newRequestWithFile(ctx, http.MethodPost, "/_api/attachments.add", params, &at, file)
	if err != nil {
		return nil, err
	}
	return &at, nil
}

// List shows attachment list of the page. This request requires page_id
func (s *AttachmentsService) List(ctx context.Context, id string) (*Attachments, error) {
	var at Attachments
	params := url.Values{}
	params.Set("access_token", s.client.Token)
	params.Set("page_id", id)
	err := s.client.newRequest(ctx, http.MethodGet, "/_api/attachments.list", params, &at)
	if err != nil {
		return nil, err
	}
	return &at, nil
}

type Attachment struct {
	Attachment AttachmentInfo `json:"attachment"`
	Page       interface{}    `json:"page"`
	OK         bool           `json:"ok"`
}

type AttachmentInfo struct {
	FileFormat   string      `json:"fileFormat"`
	FileName     string      `json:"fileName"`
	OriginalName string      `json:"originalName"`
	FilePath     string      `json:"filePath"`
	Creator      interface{} `json:"creator"`
	ID           string      `json:"_id"`
	CreatedAt    time.Time   `json:"createdAt"`
	PageCreated  bool        `json:"pageCreated"`
	URL          string      `json:"url"`
	FileSize     int         `json:"fileSize"`
	V            int         `json:"__v"`
}

type Attachments struct {
	Attachments []AttachmentInfo `json:"attachments"`
	OK          bool             `json:"ok"`
}

package crowi

import (
	"time"
)

type Crowi struct {
	Pages       Pages
	Attachments Attachments
}

type Pages struct {
	Page       PagesPage  `json:"page"`
	Attachment Attachment `json:"attachment"`
	OK         bool       `json:"ok"`
	Error      string     `json:"error"`
}

type Attachments struct {
	Page       AttachmentsPage `json:"page"`
	Attachment Attachment      `json:"attachment"`
	Filename   string          `json:"filename"`
	OK         bool            `json:"ok"`
	Error      string          `json:"error"`
}

type PagesPage struct {
	Revision       PagesRevision  `json:"revision"`
	_ID            string         `json:"_id"`
	RedirectTo     interface{}    `json:"redirectTo"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	LastUpdateUser LastUpdateUser `json:"lastUpdateUser"`
	Creator        Creator        `json:"creator"`
	Path           string         `json:"path"`
	__V            int            `json:"__v"`
	CreatedAt      time.Time      `json:"createdAt"`
	CommentCount   int            `json:"commentCount"`
	SeenUsers      []string       `json:"seenUsers"`
	Liker          []interface{}  `json:"liker"`
	GrantedUsers   []string       `json:"grantedUsers"`
	Grant          int            `json:"grant"`
	Status         string         `json:"status"`
	ID             string         `json:"id"`
}

type AttachmentsPage struct {
	Revision       AttachmentsRevision `json:"revision"`
	_ID            string              `json:"_id"`
	RedirectTo     interface{}         `json:"redirectTo"`
	UpdatedAt      time.Time           `json:"updatedAt"`
	LastUpdateUser LastUpdateUser      `json:"lastUpdateUser"`
	Creator        Creator             `json:"creator"`
	Path           string              `json:"path"`
	__V            int                 `json:"__v"`
	CreatedAt      time.Time           `json:"createdAt"`
	CommentCount   int                 `json:"commentCount"`
	SeenUsers      []string            `json:"seenUsers"`
	Liker          []interface{}       `json:"liker"`
	GrantedUsers   []string            `json:"grantedUsers"`
	Grant          int                 `json:"grant"`
	Status         string              `json:"status"`
	ID             string              `json:"id"`
}

type Attachment struct {
	__V          int       `json:"__v"`
	FileFormat   string    `json:"fileFormat"`
	FileName     string    `json:"fileName"`
	OriginalName string    `json:"originalName"`
	FilePath     string    `json:"filePath"`
	Creator      string    `json:"creator"`
	Page         string    `json:"page"`
	_ID          string    `json:"_id"`
	CreatedAt    time.Time `json:"createdAt"`
	FileSize     int       `json:"fileSize"`
}

type PagesRevision struct {
	__V       int       `json:"__v"`
	Author    string    `json:"author"`
	Body      string    `json:"body"`
	Path      string    `json:"path"`
	_ID       string    `json:"_id"`
	CreatedAt time.Time `json:"createdAt"`
	Format    string    `json:"format"`
}

type AttachmentsRevision struct {
	_ID    string `json:"_id"`
	Author struct {
		_ID       string    `json:"_id"`
		Email     string    `json:"email"`
		Username  string    `json:"username"`
		Name      string    `json:"name"`
		Admin     bool      `json:"admin"`
		CreatedAt time.Time `json:"createdAt"`
		Status    int       `json:"status"`
	} `json:"author"`
	Body      string    `json:"body"`
	Path      string    `json:"path"`
	__V       int       `json:"__v"`
	CreatedAt time.Time `json:"createdAt"`
	Format    string    `json:"format"`
}

type LastUpdateUser struct {
	_ID       string    `json:"_id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	Admin     bool      `json:"admin"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

type Creator struct {
	_ID       string    `json:"_id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

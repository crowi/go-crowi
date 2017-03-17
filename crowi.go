package crowi

import (
	"time"
)

// Crowi represents generic api json response
type Crowi struct {
	Attachment Attachment `json:"attachment"`
	Error      string     `json:"error"`
	Filename   string     `json:"filename"`
	OK         bool       `json:"ok"`
	Page       Page       `json:"page"`
}

type Page struct {
	CommentCount   int            `json:"commentCount"`
	CreatedAt      time.Time      `json:"createdAt"`
	Creator        Creator        `json:"creator"`
	Grant          int            `json:"grant"`
	GrantedUsers   []string       `json:"grantedUsers"`
	ID             string         `json:"id"`
	LastUpdateUser LastUpdateUser `json:"lastUpdateUser"`
	Liker          []interface{}  `json:"liker"`
	Path           string         `json:"path"`
	RedirectTo     interface{}    `json:"redirectTo"`
	Revision       Revision       `json:"revision"`
	SeenUsers      []string       `json:"seenUsers"`
	Status         string         `json:"status"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	_ID            string         `json:"_id"`
	__V            int            `json:"__v"`
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

type Revision struct {
	Author    interface{} `json:"author"`
	Body      string      `json:"body"`
	CreatedAt time.Time   `json:"createdAt"`
	Format    string      `json:"format"`
	Path      string      `json:"path"`
	_ID       string      `json:"_id"`
	__V       int         `json:"__v"`
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

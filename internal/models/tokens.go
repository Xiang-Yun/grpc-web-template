package models

// Token is the type for authentication tokens
type Token struct {
	PlainText string `json:"token"`
	UserID    int64  `json:"-"`
	Hash      []byte `json:"-"`
	Expiry    string `json:"expiry"`
	Scope     string `json:"-"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
}

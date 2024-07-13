package models

// User is the type for users
type User struct {
	ID        int    `json:"id"`
	User      string `json:"user"`
	Password  string `json:"password"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"updatedAt"`
}

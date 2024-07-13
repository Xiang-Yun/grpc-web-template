package repository

import (
	"grpc-web-template/internal/models"
)

// Repository is the interface which must be satisfied in order to
// connect to a database
type Repository interface {
	Migrate() error

	// users
	GetUserByUser(user string) (models.User, error)
	GetAllUsers() ([]*models.User, error)
	GetOneUser(id int) (models.User, error)
	EditUser(u models.User) error
	AddUser(u models.User, hash string) error
	DeleteUser(id int) error
	UpdatePasswordForUser(u models.User, hash string) error

	// tokens
	GetUserForToken(token string) (*models.User, error)
	AddToken(t *models.Token, u models.User) error
	Authenticate(user, password string) (int, error)
}

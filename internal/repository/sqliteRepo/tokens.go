package sqliteRepo

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"errors"
	"grpc-web-template/internal/common"
	"grpc-web-template/internal/models"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	ScopeAuthentication = "authentication"
)

// GenerateToken generates a token that lasts for ttl, and returns it
func GenerateToken(userID int, ttl time.Duration, scope string) (*models.Token, error) {
	token := &models.Token{
		UserID: int64(userID),
		Expiry: time.Now().Add(ttl).Format(common.TimeFormat),
		Scope:  scope,
	}

	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	token.PlainText = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)
	hash := sha256.Sum256(([]byte(token.PlainText)))
	token.Hash = hash[:]
	return token, nil
}

func (m *SQLiteRepository) AddToken(t *models.Token, u models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// delete existing tokens
	stmt := `delete from tokens where user_id = ?`
	_, err := m.DB.ExecContext(ctx, stmt, u.ID)
	if err != nil {
		return err
	}

	stmt = `insert into tokens (user_id, user, token_hash, expiry, createdAt, updatedAt)
			values (?, ?, ?, ?, ?, ?)`

	_, err = m.DB.ExecContext(ctx, stmt,
		u.ID,
		u.User,
		t.Hash,
		t.Expiry,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}
	return nil
}

func (m *SQLiteRepository) GetUserForToken(token string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tokenHash := sha256.Sum256([]byte(token))
	var user models.User

	query := `
		select
			u.id, u.user
		from
			users u
			inner join tokens t on (u.id = t.user_id)
		where
			t.token_hash = ?
			and t.expiry > ?
	`

	err := m.DB.QueryRowContext(ctx, query, tokenHash[:], time.Now()).Scan(
		&user.ID,
		&user.User,
	)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &user, nil
}

// Authenticate attempts to log a user in by comparing supplied password with password hash
func (m *SQLiteRepository) Authenticate(user, password string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id int
	var hashedPassword string

	row := m.DB.QueryRowContext(ctx, "select id, password from users where user = ?", user)
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		return id, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, errors.New("incorrect password")
	} else if err != nil {
		return 0, err
	}

	return id, nil
}

package sqliteRepo

import (
	"context"
	"grpc-web-template/internal/common"
	"grpc-web-template/internal/models"
	"strings"
	"time"
)

// GetUserByEmail gets a user by email address
func (repo *SQLiteRepository) GetUserByUser(user string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	user = strings.ToLower(user)
	var u models.User

	query := `
		select id, user, password, createdAt, updatedAt from users where user = ?
	`

	row := repo.DB.QueryRowContext(ctx, query, user)

	err := row.Scan(
		&u.ID,
		&u.User,
		&u.Password,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		return u, err
	}

	return u, nil
}

// GetAllUsers returns a slice of all users
func (repo *SQLiteRepository) GetAllUsers() ([]*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var users []*models.User

	query := `
		select id, user, createdAt, updatedAt from users
	`

	rows, err := repo.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u models.User
		err = rows.Scan(
			&u.ID,
			&u.User,
			&u.CreatedAt,
			&u.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &u)
	}

	return users, nil
}

// GetOneUser returns one user by id
func (repo *SQLiteRepository) GetOneUser(id int) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var u models.User

	query := `
		select
			id, user, createdAt, updatedAt
		from
			users
		where id = ?`

	row := repo.DB.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&u.ID,
		&u.User,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		return u, err
	}
	return u, nil
}

// EditUser edits an existing user
func (repo *SQLiteRepository) EditUser(u models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
		update users set
			user = ?,
			updatedAt = ?
		where
			id = ?`

	_, err := repo.DB.ExecContext(ctx, stmt,
		u.User,
		time.Now().Format(common.TimeFormat),
		u.ID,
	)

	if err != nil {
		return err
	}
	return nil
}

// AddUser inserts a user into the database
func (repo *SQLiteRepository) AddUser(u models.User, hash string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
		insert into users (user, password, createdAt, updatedAt)
		values (?, ?, ?, ?)`

	_, err := repo.DB.ExecContext(ctx, stmt,
		u.User,
		hash,
		time.Now().Format(common.TimeFormat),
		time.Now().Format(common.TimeFormat),
	)

	if err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes a user by id
func (repo *SQLiteRepository) DeleteUser(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `delete from users where id = ?`

	_, err := repo.DB.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	stmt = "delete from tokens where user_id = ?"
	_, err = repo.DB.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil
}

// UpdatePasswordForUser updates the password hash for a given user, by user id
func (repo *SQLiteRepository) UpdatePasswordForUser(u models.User, hash string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `update users set password = ? where id = ?`
	_, err := repo.DB.ExecContext(ctx, stmt, hash, u.ID)
	if err != nil {
		return err
	}

	return nil
}

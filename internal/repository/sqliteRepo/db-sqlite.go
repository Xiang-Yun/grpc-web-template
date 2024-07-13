package sqliteRepo

import "database/sql"

// SQLiteRepository the type for a repository that connects to sqlite database
type SQLiteRepository struct {
	DB *sql.DB
}

// NewSQLiteRepository returns a new repository with a connection to sqlite
func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		DB: db,
	}
}

// Migrate creates the table(s) we need
func (repo *SQLiteRepository) Migrate() error {
	// Create users table
	userQuery := `
		create table if not exists users (
			id integer primary key autoincrement,
			user TINYTEXT,
			password TEXT,
			createdAt TINYTEXT,
			updatedAt TINYTEXT);
		`
	_, err := repo.DB.Exec(userQuery)
	if err != nil {
		return err
	}

	// Create tokens table
	tokenQuery := `
		create table if not exists tokens (
			id integer primary key autoincrement,
			user_id INTEGER,
			user TINYTEXT,
			token_hash TEXT,
			expiry TINYTEXT,
			createdAt TINYTEXT,
			updatedAt TINYTEXT);
		`
	_, err = repo.DB.Exec(tokenQuery)
	if err != nil {
		return err
	}

	// Create storage session table
	sessionQuery := `
		create table if not exists sessions (
			token TEXT primary key,
			data BLOB,
			expiry REAL NOT NULL);
		`
	_, err = repo.DB.Exec(sessionQuery)
	if err != nil {
		return err
	}

	// insert one user for default
	oneUser := `
		INSERT INTO users (user, password, createdAt, updatedAt)
			SELECT 'admin', 
				'$2a$10$u2KDqyKJBYy3HM79DEBBleq/ZscFS6ybRkLNnCoAquSX7vh/SmP/G',
				strftime('%Y-%m-%d %H:%M:%S.000', datetime('now', '+8 hours')),
				strftime('%Y-%m-%d %H:%M:%S.000', datetime('now', '+8 hours'))
			WHERE NOT EXISTS (SELECT 1 FROM users);
	`
	_, err = repo.DB.Exec(oneUser)
	if err != nil {
		return err
	}

	return nil
}

func ConnectSQL(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func SetupDB(sqlDB *sql.DB) *SQLiteRepository {
	db := NewSQLiteRepository(sqlDB)
	err := db.Migrate()
	if err != nil {
		panic(err)
	}
	return db
}

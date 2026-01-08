package models

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/T2Knock/snippetbox/internal/auth"
	"github.com/go-sql-driver/mysql"
)

type User struct {
	ID              int
	NAME            string
	EMAIL           string
	HASHED_PASSWORD []byte
	CREATED         time.Time
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email, password string) error {
	hashed_password, err := auth.HashPassword(password)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (name, email, hashed_password, created)
    VALUES(?, ?, ?,UTC_TIMESTAMP())`

	if _, err := m.DB.Exec(stmt, name, email, hashed_password); err != nil {
		var mysqlError *mysql.MySQLError
		if errors.As(err, &mysqlError) {
			if mysqlError.Number == 1062 && strings.Contains(mysqlError.Message, "users_uc_email") {
				return ErrDuplicateEmail
			}
		}
	}

	return nil
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

func (m *UserModel) Exist(id int) (bool, error) {
	stmt := `SELECT id WHERE id = ?`

	s := &Snippet{}
	if err := m.DB.QueryRow(stmt, id).Scan(&s.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return true, nil
	}

	return false, nil
}

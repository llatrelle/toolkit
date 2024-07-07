package utils

import (
	"database/sql"
	"fmt"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func NewMockDB() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Printf("Error creating sqlMock: %v", err.Error())
	}
	return db, mock
}

package tests

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Printf("Error creating sqlMock: %v", err.Error())
	}
	return db, mock
}

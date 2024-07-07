package tests

import (
	"database/sql"
	"fmt"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Printf("Error creating sqlMock: %v", err.Error())
	}
	return db, mock
}

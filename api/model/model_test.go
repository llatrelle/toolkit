package model

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestGetVersionStringDBNil(t *testing.T) {
	db = nil
	version, err := getSQLVersion()
	assert.NoError(t, err)
	assert.Equal(t, "", version)
}

func TestGetVersionStringDBWithError(t *testing.T) {

	version, err := getSQLVersion()
	assert.NoError(t, err)
	assert.Equal(t, "", version)
}

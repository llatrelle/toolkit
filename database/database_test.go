package database

import (
	"database/sql"
	"fmt"
	"toolkit/tests/utils"

	"github.com/stretchr/testify/assert"

	"testing"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Print("Error")
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

func TestGetVersionStringDBSuccess(t *testing.T) {
	versionExpected := "8.0.28"
	dbcon, mock := utils.NewMockDB()
	db = dbcon
	rows := sqlmock.NewRows([]string{"VERSION()"}).AddRow(versionExpected)
	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(rows)
	version, err := getSQLVersion()
	assert.NoError(t, err)
	assert.Equal(t, versionExpected, version)
}

func TestVerifyConnectionSettingsUserError(t *testing.T) {
	err := verifyConnectionSettings("", "pass", "localhost", "core")
	assert.Equal(t, "invalid SQL user", err.Error())
}

func TestVerifyConnectionSettingsSecretError(t *testing.T) {
	err := verifyConnectionSettings("root", "", "localhost", "core")
	assert.Equal(t, "invalid SQL secret", err.Error())
}

func TestVerifyConnectionSettingsHostError(t *testing.T) {
	err := verifyConnectionSettings("root", "1234", "", "core")
	assert.Equal(t, "invalid SQL host", err.Error())
}

func TestVerifyConnectionSettingsSchemaError(t *testing.T) {
	err := verifyConnectionSettings("root", "1234", "localhost", "")
	assert.Equal(t, "invalid SQL schema", err.Error())
}

func TestVerifyConnectionSettingsSchemaSuccess(t *testing.T) {
	err := verifyConnectionSettings("root", "1234", "localhost", "core")
	assert.NoError(t, err)
}

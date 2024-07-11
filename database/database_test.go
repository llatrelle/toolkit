package database

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/llatrelle/toolkit/tests/utils"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestGetVersionStringDBNil(t *testing.T) {
	db = nil
	version, err := getSQLVersion()
	assert.Equal(t, err.Error(), ErrorSQLConnection.Error())
	assert.Equal(t, "", version)
}

func TestGetVersionStringDBSuccess(t *testing.T) {
	versionExpected := "8.0.28"
	gdb, mock := utils.NewMockDB()
	db = gdb
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

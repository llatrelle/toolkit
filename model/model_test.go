package model

import (
	tests "github.com/llatrelle/toolkit/tests/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestAddDBNilPointerError(t *testing.T) {
	defer cleanDatabases()

	err := AddDB("test-core", nil)
	assert.Equal(t, "connection can not be nil", err.Error())

}

func TestAddDBConnectionNameError(t *testing.T) {
	defer cleanDatabases()

	db, _ := tests.NewMockDB()
	err := AddDB("", db)
	assert.Equal(t, "connection name can not be empty", err.Error())
}

func TestAddDBExistConnectionNameError(t *testing.T) {
	defer cleanDatabases()

	db, _ := tests.NewMockDB()
	err := AddDB("test-core", db)
	assert.NoError(t, err)
	err = AddDB("test-core", db)
	assert.Equal(t, "Error adding database, connection exist", err.Error())

}

func TestAddDBConnectionSuccess(t *testing.T) {
	defer cleanDatabases()

	connName := "test-core"
	db, _ := tests.NewMockDB()
	err := AddDB(connName, db)
	assert.NoError(t, err)
	sameDB := GetDB(connName)
	assert.NotEqual(t, nil, sameDB)
	assert.Equal(t, &db, &sameDB)

}

func TestRemoveDBNotExistError(t *testing.T) {
	defer cleanDatabases()

	connName := "test-core"
	connNameNotExist := "test"
	gdb, _ := tests.NewMockDB()
	err := AddDB(connName, gdb)
	assert.NoError(t, err)
	err = RemoveDB(connNameNotExist)
	assert.Equal(t, "the connection does not exist", err.Error())
	assert.Equal(t, 1, len(databases))
}

func TestRemoveDBSuccess(t *testing.T) {
	defer cleanDatabases()

	connName := "test-core"
	gdb, _ := tests.NewMockDB()
	db, err := gdb.DB()
	assert.NoError(t, db.Ping())
	err = AddDB(connName, gdb)
	assert.NoError(t, err)
	err = RemoveDB(connName)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(databases))
}

func cleanDatabases() {
	databases = make(map[string]*gorm.DB)
}

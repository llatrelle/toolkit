package database

import (
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type dbConnector interface {
	Open(string, string, string, string) *gorm.DB
}

// Connect Open a SQL connection
func Connect(typeDB, user, secret, server, schema string) (*gorm.DB, error) {

	if err := verifyConnectionSettings(user, secret, server, schema); err != nil {
		return nil, err
	}
	if db != nil {
		version, err := getSQLVersion()
		if err != nil || version == "" {
			return db, err
		}
		return db, nil
	}
	dbconn := newDB(typeDB)
	db = dbconn.Open(user, secret, server, schema)

	return db, nil
}

// newDB Get a DBConnector with specific driver
func newDB(typeDB string) dbConnector {
	switch typeDB {
	case DBTypeMySQL:
		dbType := mysqlConnector{}
		return dbType
		break
	default:
		dbType := genericDB{}
		return dbType
	}
	return nil
}

type genericDB struct{}

func (m genericDB) Open(user, secret, server, schema string) *gorm.DB {
	return nil
}

// getSQLVersion return de SQL version on the server
func getSQLVersion() (string, error) {
	var version string
	conn, err := db.DB()
	if err == nil && db != nil {
		err := conn.QueryRow("SELECT VERSION()").Scan(&version)
		return version, err

	}
	return version, nil
}

func verifyConnectionSettings(user, secret, server, schema string) error {
	switch {
	case user == "":
		return ErrorSQLUser
	case secret == "":
		return ErrorSQLSecret
	case server == "":
		return ErrorSQLHost
	case schema == "":
		return ErrorSQLSchema
	}

	return nil

}

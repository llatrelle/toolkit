package model

import (
	"database/sql"
	"errors"
	"github.com/llatrelle/toolkit/config"
)

var db *sql.DB

type dbConnector interface {
	Open(string, string, string, string) (*sql.DB, error)
}

//Connect Open a SQL connection
func Connect(typeDB, user, secret, server, schema string) (*sql.DB, error) {

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
	db, err := dbconn.Open(user, secret, server, schema)
	if err != nil {
		return nil, err
	}
	return db, err
}

//newDB Get a DBConnector with specific driver
func newDB(typeDB string) dbConnector {
	switch typeDB {
	case config.DBTypeMySQL:
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

func (m genericDB) Open(user, secret, server, schema string) (*sql.DB, error) {
	return nil, errors.New("driver is not implemented")
}

//getSQLVersion return de SQL version on the server
func getSQLVersion() (string, error) {
	var version string
	if db != nil {
		err := db.QueryRow("SELECT VERSION()").Scan(&version)
		return version, err

	}
	return version, nil
}

func verifyConnectionSettings(user, secret, server, schema string) error {
	switch {
	case user == "":
		return errors.New(config.ErrorSQLUser)
	case secret == "":
		return errors.New(config.ErrorSQLSecret)
	case server == "":
		return errors.New(config.ErrorSQLHost)
	case schema == "":
		return errors.New(config.ErrorSQLSchema)
	}

	return nil

}

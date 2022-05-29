package database

import (
	"database/sql"
	"errors"
	"github.com/llatrelle/toolkit/config"
	"github.com/rs/zerolog"
)

var (
	log zerolog.Logger
	db  *sql.DB
)

type dbConnector interface {
	Open(string, string, string, string) *sql.DB
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
	db = dbconn.Open(user, secret, server, schema)

	return db, nil
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

func (m genericDB) Open(user, secret, server, schema string) *sql.DB {
	return nil
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

func SetLogger(l zerolog.Logger) {
	log = l
}

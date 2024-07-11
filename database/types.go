package database

import "errors"

const (
	DBTypeMySQL    = "mysql"
	DBTypePostgres = "postgres"
	DBPrimary      = "DB_PRIMARY"
)

var (
	ErrorSQLUser       = errors.New("invalid SQL user")
	ErrorSQLSecret     = errors.New("invalid SQL secret")
	ErrorSQLHost       = errors.New("invalid SQL host")
	ErrorSQLSchema     = errors.New("invalid SQL schema")
	ErrorSQLConnection = errors.New("Error nil connection")
)

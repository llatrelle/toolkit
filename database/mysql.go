package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func newMySQLConnection(user, password, schema, host string) (*sql.DB, error) {
	conn, err := sql.Open("mysql", user+":"+password+"@tcp("+host+")/"+schema)
	if err != nil {
		return nil, err
	}
	err = conn.Ping()
	if err != nil {
		return nil, err
	}
	return conn, err
}

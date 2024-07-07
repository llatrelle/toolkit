package database

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
)

func newMySQLConnection(user, password, schema, host string) (*sql.DB, error) {

	dsn := fmt.Sprintf("%s:%S@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, schema)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	err = conn.Ping()
	if err != nil {
		return nil, err
	}
	return conn, err
}

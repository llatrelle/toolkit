package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/llatrelle/toolkit/config"
	"time"
)

type mysqlConnector struct{}

func (m mysqlConnector) Open(user, secret, server, schema string) (*sql.DB, error) {
	connString := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, secret, server, schema)
	for {
		db, err := sql.Open(config.DBTypeMySQL, connString)
		if err == nil {
			return db, err
		}
		log.Error().Err(err).Msg("Error connecting to MySQL Server (%v)")
		time.Sleep(time.Second * 3)
	}

}

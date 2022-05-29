package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/llatrelle/toolkit/config"
	"time"
)

type mysqlConnector struct{}

func (m mysqlConnector) Open(user, secret, server, schema string) *sql.DB {
	connString := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, secret, server, schema)
	for {
		db, err := sql.Open(config.DBTypeMySQL, connString)
		if err == nil {
			if err = db.Ping(); err == nil {
				return db
			}
		}
		log.Error().Err(err).Msgf("Error connecting to MySQL Server (%v)", server)
		time.Sleep(time.Second * 3)
	}

}

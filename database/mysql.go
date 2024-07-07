package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"toolkit/logger"
)

type mysqlConnector struct{}

func (m mysqlConnector) Open(user, secret, server, schema string) *gorm.DB {
	connString := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, secret, server, schema)
	for {
		gormDB, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
		//	db, err := sql.Open(config.DBTypeMySQL, connString)
		if err == nil {
			conn, err =: gormDB.DB()
			if err == nil && conn.Ping() == nil {
				return db
			}

		}
		logger.Error("Error connecting to MySQL Server", nil)
		time.Sleep(time.Second * 3)
	}

}

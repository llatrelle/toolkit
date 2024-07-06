package db

import (
	"database/sql"
	"fmt"
	"time"
	"toolkit/logger"
)

var db *sql.DB

type DB struct {
	connection *sql.DB
}

type modeler interface {
	Get(int) error
}

func SetConnection(user, password, schema, host string) error {
	if db != nil {
		return nil
	}

	var err error

	for {
		db, err = newMySQLConnection(user, password, schema, host)
		if err == nil {
			logger.Info("Database connected successfully")
			break
		}

		logger.Error("Error connecting to database... ", err)
		time.Sleep(time.Second * 3)
	}

	return nil

}

func GetSQLQuery(sql string, params interface{}) ([]map[string]interface{}, error) {
	rows, err := db.Query(sql, params)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	cols, _ := rows.Columns()
	var r []map[string]interface{}
	var values []interface{}
	for rows.Next() {
		values = make([]interface{}, len(cols))
		for i := range values {
			values[i] = &values[i]
		}
		err = rows.Scan(values...)
		if err != nil {
			return nil, err
		}
		var row Result
		row = make(Result)
		for i, col := range cols {
			logger.Debug(fmt.Sprintf("%T", values[i]))
			row[col] = values[i]
		}
		r = append(r, row)
	}

	return r, err
}

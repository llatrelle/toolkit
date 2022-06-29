package model

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

var (
	databases map[string]*sql.DB
)

func init() {
	databases = make(map[string]*sql.DB)
}

//AddDB Add a conection in map, with a connection name
func AddDB(connectionName string, db *sql.DB) error {
	if connectionName == "" {
		return errors.New("connection name can not be empty")
	}
	if db == nil {
		return errors.New("connection can not be nil")
	}
	if _, v := databases[connectionName]; v {
		return errors.New("Error adding database, connection exist")
	}
	databases[connectionName] = db
	return nil
}

//GetDB Return a *sql.DB
func GetDB(connectionName string) *sql.DB {
	return databases[connectionName]
}

//RemoveDB Remove a connection from map of connection. Befre delete, try close the connection
func RemoveDB(connectionName string) error {
	conn, hasValue := databases[connectionName]
	if !hasValue || conn == nil {
		return errors.New("the connection does not exist")
	}
	if err := conn.Close(); err != nil && err.Error() != "all expectations were already fulfilled, call to database Close was not expected" {
		return err
	}
	delete(databases, connectionName)
	return nil
}

type Modeler interface {
	GetKeys() []string
	TableName() string
	Scan(*sql.Rows) Modeler
}

func Get(m *Modeler, pks []string) error {
	db := GetDB("primary")

	if keys := getKeyPair((*m).GetKeys(), pks); keys != "" {
		q := fmt.Sprintf("SELECT * FROM %s WHERE %s", (*m).TableName(), keys)
		row, err := db.Query(q)
		if err != nil {
			return err
		}
		defer row.Close()
		if row.Next() {
			*m = (*m).Scan(row)
		}

	} else {
		return errors.New("invalid keys")
	}
	return nil
}

func Delete(m Modeler, pks []string) (int64, error) {
	db := GetDB("primary")

	if keys := getKeyPair(m.GetKeys(), pks); keys != "" {
		q := fmt.Sprintf("DELETE FROM %s WHERE %s", m.TableName(), keys)
		result, err := db.Exec(q)
		if err != nil {
			return 0, err
		}
		code, err := result.RowsAffected()
		if err != nil {
			return 0, err
		}
		return code, nil

	} else {
		return 0, errors.New("invalid keys")
	}

}
func Create(m Modeler, payload map[string]interface{}) {
	db := GetDB("primary")

	q := fmt.Sprintf("Insert * from %s;", m.TableName())
}
func GetAll(m Modeler, result *[]Modeler) error {
	db := GetDB("primary")

	//Need to reeplace arguments
	q := fmt.Sprintf("select * from %s;", m.TableName())
	rows, err := db.Query(q)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		newModel := m.Scan(rows)
		*result = append(*result, newModel)
	}
	return err
}

func getKeyPair(keys, pks []string) string {
	if len(keys) != len(pks) {
		return ""
	}
	var listKeys []string
	for i := range keys {
		keypair := fmt.Sprintf("%s=%s", keys[i], pks[i])
		listKeys = append(listKeys, keypair)
	}
	return strings.Join(listKeys, " AND ")
}

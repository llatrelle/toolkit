package repository

import (
	"errors"
	"fmt"
	"github.com/llatrelle/toolkit/logger"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

//TODO:this package not should be return http status code

const (
	// DBPrimary TODO: Move this const to a config file or package
	DBPrimary = "DBPrimary"
)

type repository struct {
	databases map[string]*gorm.DB
}

// NewRepository Create a new repository
func NewRepository() *repository {
	return &repository{databases: make(map[string]*gorm.DB)}
}

// AddDB Add a conection in map, with a connection name
func (r *repository) AddDB(connectionName string, db *gorm.DB) error {
	if connectionName == "" {
		return errors.New("connection name can not be empty")
	}
	if db == nil {
		return errors.New("connection can not be nil")
	}
	if _, v := r.databases[connectionName]; v {
		return errors.New("Error adding database, connection exist")
	}

	r.databases[connectionName] = db
	return nil
}

// GetDB Return a *sql.DB
func (r *repository) GetDB(connectionName string) *gorm.DB {
	return r.databases[connectionName]
}

// RemoveDB Remove a connection from map of connection. Befre delete, try close the connection
func (r *repository) RemoveDB(connectionName string) error {
	db, hasValue := r.databases[connectionName]
	if !hasValue {
		return errors.New("the connection does not exist")
	}
	conn, _ := db.DB()
	if conn == nil {
		return errors.New("the connection does not exist")
	}
	if err := conn.Close(); err != nil && err.Error() != "all expectations were already fulfilled, call to database Close was not expected" {
		return err
	}
	delete(r.databases, connectionName)
	return nil
}

func (r *repository) Register(modeler Modeler) {
	//TODO: use const for primary connectionName
	db := r.GetDB(DBPrimary)
	logger.Info(fmt.Sprintf("Automigrando recurso %v ... ", modeler.TableName()))
	if err := db.AutoMigrate(modeler); err != nil {
		logger.Error(fmt.Sprintf("Error automigrando recurso %v ", modeler.TableName()), err)
	}
}

type Modeler interface {
	GetKeys() []string
	TableName() string
	SetPKs([]string)
	NewModel() Modeler
	NewModelList() interface{}
}

func (r *repository) Get(m *Modeler) (int, error) {
	db := r.GetDB(DBPrimary)
	var rowsLen int64
	rowsLen = 0

	err := db.Find(m, *m).Count(&rowsLen).Set("gorm:auto_preload", true).Error
	if err != nil {
		return http.StatusInternalServerError, errors.New("Error consultado base de datos")
	}
	if rowsLen > 1 {
		return http.StatusConflict, errors.New("expect 1 result and get more.")
	}
	if rowsLen == 0 {
		return http.StatusNotFound, errors.New("resource not found")
	}
	return http.StatusOK, nil
}

func (r *repository) Delete(m Modeler) (int, error) {
	db := r.GetDB(DBPrimary)
	scope := db.Delete(m)
	if scope.Error != nil {
		return http.StatusInternalServerError, errors.New("Error deleting the resource")
	}
	if scope.RowsAffected == 0 {
		return http.StatusNotFound, errors.New("Resource not found")
	}
	if scope.RowsAffected > 1 {
		return http.StatusConflict, errors.New("Many resources are deleted")
	}
	return http.StatusOK, nil
}
func (r *repository) Create(m Modeler) (int, error) {
	db := r.GetDB(DBPrimary)
	err := db.Create(m).Error
	if err != nil {
		return http.StatusInternalServerError, errors.New("Error creating resource")
	}
	return http.StatusCreated, err
}

func (r *repository) Update(m Modeler) (int, error) {
	db := r.GetDB(DBPrimary)
	scope := db.Model(m).Updates(m)

	if scope.Error != nil {
		return http.StatusInternalServerError, errors.New("Error updating resource")
	}
	if scope.RowsAffected == 0 {
		return http.StatusNotFound, errors.New("Resource not found")
	}
	if scope.RowsAffected > 1 {
		scope.Rollback()
		return http.StatusConflict, errors.New("Warning!,many resources updated. Rollbacked! ")
	}
	return http.StatusOK, nil
}

func (r *repository) GetAll(m Modeler, result interface{}) error {
	db := r.GetDB(DBPrimary)
	err := db.Model(m).Find(result).Error
	return err
}

func (r *repository) getKeyPair(keys, pks []string) string {
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

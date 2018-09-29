package models

import (
	"database/sql"
	"fmt"
	// Import mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/juju/errors"
)

var (
	globalDB *sql.DB
)

type rowScanner interface {
	Scan(...interface{}) error
}

type DataSourceConfig struct {
	IP       string `toml:"ip"`
	Port     int    `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Database string `toml:"database"`
}

func (d *DataSourceConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		d.Username, d.Password, d.IP, d.Port, d.Database)
}

func InitGlobalDB(dsn *DataSourceConfig) error {
	var err error
	globalDB, err = sql.Open("mysql", dsn.DSN())
	if nil != err {
		return errors.Trace(err)
	}
	/*if err = globalDB.Ping(); nil != err {
		return errors.Trace(err)
	}*/
	return nil
}

func GetGlobalDB() *sql.DB {
	return globalDB
}

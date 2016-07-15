package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/wagnergaldino/training/dne/util"
)

func NewDB() *sql.DB {
	var err error
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/dne?charset=utf8")
	util.CheckErr(err)
	return db
}

var db *sql.DB

func Init() {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/dne?charset=utf8")
	util.CheckErr(err)
}

func Get() *sql.DB {
	return db
}

func Close() error {
	return db.Close()
}

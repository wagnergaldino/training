package orm

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"

	"github.com/wagnergaldino/training/dne/util"
)

var db *gorm.DB

func Init() {
	var err error
	if db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/dne?charset=utf8"); err == nil {
		if err = db.DB().Ping(); err == nil {
			db.LogMode(true)
			db.SingularTable(true)
		}
	}
	util.CheckErr(err)
}

func Get() *gorm.DB {
	return db
}

func Close() error {
	return db.DB().Close()
}

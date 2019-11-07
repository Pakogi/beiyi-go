package model

import (
	"time"

	"beiyi/util"

	"github.com/jinzhu/gorm"

	//
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func Database(connString string) {
	db, err := gorm.Open("postgres", connString)

	if err != nil {
		util.Log().Panic("Database connection failure!", err)
	}

	db.LogMode(true)

	db.DB().SetMaxIdleConns(50)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db
}

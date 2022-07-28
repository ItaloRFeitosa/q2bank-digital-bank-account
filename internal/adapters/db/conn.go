package db

import (
	"github.com/italorfeitosa/q2bank-digital-bank-account/internal/shared/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.DBUrl), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.Debug().AutoMigrate(&User{}, &Account{}, &Transfer{}, &Transaction{})
	if err != nil {
		panic("failed to migrate models")
	}
	return db
}

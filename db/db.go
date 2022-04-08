package db

import (
	"demo/types"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	dbpath := os.Getenv("DB_PATH")
	if dbpath == "" {
		dbpath = "screech.db"
	}

	log.Printf("db file: %v", dbpath)
	db, err := gorm.Open(sqlite.Open(dbpath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&types.User{}, &types.Screech{})
	if err != nil {
		panic("failed to migrate schema")
	}

	DB = db
}

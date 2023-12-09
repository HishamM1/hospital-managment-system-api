package config

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectDB connects go to mysql database
func ConnectDB() *gorm.DB {
	db, errorDB := gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{})
	if errorDB != nil {
		panic("Failed to connect mysql database")
	}

	return db
}

// DisconnectDB is stopping your connection to mysql database
func DisconnectDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to kill connection from database")
	}
	dbSQL.Close()
}

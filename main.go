package main

import (
	"main/config"
	"main/routes"

	"gorm.io/gorm"
)

var db *gorm.DB = config.ConnectDB()

func main() {
	defer config.DisconnectDB(db)

	routes.Api()
}

package main

import (
	"gorm.io/gorm"
	"todoList/src/config"
	"todoList/src/routes"
)

var (
	db *gorm.DB = config.ConnectDB()
)

func main() {
	defer config.DisconnectDB(db)

	//run all routes
	routes.Routes()
}

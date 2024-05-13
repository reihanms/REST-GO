package main

import (
	"gin-gorm/internal/app"
	"gin-gorm/internal/database"
)

func main() {
	database.ConnectDB()
	app.Run()
}

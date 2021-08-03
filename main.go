package main

import (
	migration "Projects/go_framework_gofiber/Migration"
	service "Projects/go_framework_gofiber/Service"
)


//package installed
//go get -u github.com/gofiber/fiber/v2
// go get -u gorm.io/gorm
// go get -u gorm.io/driver/mysql  

func main() {
	migration.InitialMigration()
	service.Initialize()
}

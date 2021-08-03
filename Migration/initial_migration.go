package migration

import (
	model "Projects/go_framework_gofiber/Models"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var context *gorm.DB
var err error


const connectionString = "root:faizan123@tcp(127.0.0.1)/gofiberdb?charset=utf8mb4&parseTime=True&loc=Local"


func InitialMigration()  {
	context, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}

	context.AutoMigrate(&model.User{})
}


func GetUsers(cons *fiber.Ctx) error {
	var users []model.User
	context.Find(&users)
	return cons.JSON(&users)
}

func GetUser(cons *fiber.Ctx) error {
	id := cons.Params("id")
	var user model.User
	context.Find(&user, id)
	return cons.JSON(&user)
}

func SaveUser(cons *fiber.Ctx) error {
	user := new(model.User)
	if err := cons.BodyParser(user); err != nil {
		return cons.Status(500).SendString(err.Error())
	}
	context.Create(&user)
	return cons.JSON(&user)
}

func UpdateUser(cons *fiber.Ctx) error {
	id := cons.Params("id")
	user := new(model.User)
	context.First(&user, id)
	if user.Email == "" {
		return cons.Status(500).SendString("User not available")
	}
	if err := cons.BodyParser(user); err != nil {
		return cons.Status(500).SendString(err.Error())
	}
	context.Save(&user)
	return cons.JSON(&user)
}

func DeleteUser(cons *fiber.Ctx) error {
	id := cons.Params("id")
	var user model.User
	context.First(&user, id)
	if user.Email == "" {
		return cons.Status(500).SendString("User not available")
	}
	context.Delete(&user)
	return cons.SendString("User deleted successfully")
}




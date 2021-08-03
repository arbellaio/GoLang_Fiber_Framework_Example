package service

import (
	migration "Projects/go_framework_gofiber/Migration"
	"github.com/gofiber/fiber/v2"
)

func Initialize() {
	app := fiber.New()
	routers(app)
	app.Get("/", index)
	app.Listen(":3000")
}

func index(con *fiber.Ctx) error {
	return con.SendString("Hello Welcome to GoLang Tutorial")
}

func routers(app *fiber.App)  {
	app.Get("/api/users", migration.GetUsers)
	app.Get("/api/user/:id", migration.GetUser)
	app.Post("/api/user", migration.SaveUser)
	app.Delete("/api/user/:id", migration.DeleteUser)
	app.Put("/api/user/:id", migration.UpdateUser)
}

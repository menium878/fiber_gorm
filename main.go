package main

import (
	"fiber_gorm/controllers"
	"fiber_gorm/initializers"
	"os"

	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	app := fiber.New()

	app.Get("/posts", controllers.PostRead)
	app.Post("/post", controllers.PostCreate)
	app.Get("post/:id", controllers.PostReadOne)
	app.Put("/post/:id", controllers.PostUpdate)
	app.Delete("/post/:id", controllers.PostDelete)
	PORT := os.Getenv("PORT")
	app.Listen(":" + PORT)
}

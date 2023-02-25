package main

import (
	"fiber_gorm/controllers"
	"fiber_gorm/initializers"
	"os"

	_ "fiber_gorm/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/swagger"
)

// @title Fiber Example API
// @version 2.0
// @description My fiber app for learning
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Use(logger.New())                                                            // middleware - logger do logowania
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Monitor of the server"})) // middleware - do wyswietlania strony z statystykami strony
	app.Get("/posts", controllers.PostRead)
	app.Post("/post", controllers.PostCreate)
	app.Get("post/:id", controllers.PostReadOne)
	app.Put("/post/:id", controllers.PostUpdate)
	app.Delete("/post/:id", controllers.PostDelete)

	PORT := os.Getenv("PORT")
	app.Listen(":" + PORT)
}

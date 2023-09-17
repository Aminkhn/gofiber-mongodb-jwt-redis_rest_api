package router

import (
	handlers "github.com/aminkhn/mongo-rest-api/handlers"
	"github.com/aminkhn/mongo-rest-api/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// General Middlewares
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	// API Health Checker
	api := app.Group("/api")
	api.Get("/HealthChecker", handlers.HealthChecker)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handlers.Login)
	auth.Post("/logout", middlewares.Protected(), handlers.Logout)

	// User
	user := api.Group("/user")
	// Protection
	user.Use(middlewares.Protected())
	user.Use(middlewares.IsBlackListed)
	// User CRUD
	user.Get("/", handlers.GetAllUser)
	user.Get("/:id", handlers.GetUser)
	user.Post("/", handlers.CreateUser)
	user.Put("/:id", handlers.UpdateUserPut)
	user.Patch("/:id", handlers.UpdateUserPatch)
	user.Delete("/:id", handlers.DeleteUser)
}

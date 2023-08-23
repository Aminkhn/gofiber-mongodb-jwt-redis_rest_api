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
	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	// API Health Checker
	api := app.Group("/api", logger.New())
	api.Get("/HealthChecker", handlers.HealthChecker)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handlers.Login)
	//auth.Post("/logout", handlers.Logout)

	// User , middlewares.DeserializeUser()
	user := api.Group("/user")
	user.Get("/", middlewares.DeserializeUser, handlers.GetAllUser)
	user.Get("/:id", handlers.GetUser)
	user.Post("/", handlers.CreateUser)
	user.Put("/:id", handlers.UpdateUserPut)
	user.Patch("/:id", handlers.UpdateUserPatch)
	user.Delete("/:id", handlers.DeleteUser)
	/*
		// Product
		product := api.Group("/product")
		product.Get("/", handlers.GetAllProducts)
		product.Get("/:id", handlers.GetProduct)
		product.Post("/", middlewares.DeserializeUser(), handlers.CreateProduct)
		product.Delete("/:id", middlewares.DeserializeUser(), handlers.DeleteProduct)
		// Order

			order := api.Group("/order")
			order.Get("/", handlers.GetAllOrders)
			order.Get("/:id", handlers.GetOrder)
			order.Post("/", middlewares.DeserializeUser(), handlers.CreateOrder)
			order.Delete("/:id", middlewares.DeserializeUser(), handlers.DeleteOrder)
	*/
}

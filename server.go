package main

import (
	"log"

	"github.com/aminkhn/mongo-rest-api/config"
	"github.com/aminkhn/mongo-rest-api/database"
	"github.com/aminkhn/mongo-rest-api/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// loading Env variables
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("can not load Envirnment variables", err)
	}
	// Connect to the database
	if err := database.MongoConnect(&loadConfig); err != nil {
		log.Fatal(err)
	}
	defer database.CloseMongo()
	// connecting to Redis
	database.RedisConnectDb(&loadConfig)
	// new instance of fiber
	app := fiber.New()
	// setting up URIs routes
	router.SetupRoutes(app)
	// staring webserver
	app.Listen(":" + loadConfig.Port)
}

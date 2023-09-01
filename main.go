package main

import (
	"github.com/aminkhn/mongo-rest-api/app"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app.Server(&fiber.Ctx{})
}

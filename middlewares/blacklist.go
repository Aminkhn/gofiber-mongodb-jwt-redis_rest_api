package middlewares

import (
	"log"
	"strings"

	"github.com/aminkhn/mongo-rest-api/config"
	"github.com/aminkhn/mongo-rest-api/database"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// check if passed token exist in blacklist
func IsBlackListed(c *fiber.Ctx) error {
	reqToken := c.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) == 2 {
		reqToken = splitToken[1]
		claims := jwt.MapClaims{}
		// loading Env variables
		loadConfig, err1 := config.LoadConfig("./")
		if err1 != nil {
			log.Fatal("can not load Envirnment variables", err1)
		}
		_, err := jwt.ParseWithClaims(reqToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(loadConfig.Secret), nil
		})
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Can not Decode jwt", "data": err.Error()})
		}
		userId := claims["user_id"].(string)

		blacListed, err := database.RedisDb.Db.Get(userId).Result()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Problem with Redis reading blacklist", "data": err.Error()})
		}
		if blacListed == reqToken {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "your Token Blacklisted Login Again", "data": nil})
		}
		return c.Next()
	}
	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"status": "error", "message": "your jwt token format has problem or missing!", "data": nil})
}

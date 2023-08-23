package middlewares

import (
	"github.com/aminkhn/mongo-rest-api/database"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var UserID *redis.StringCmd

// protects routes
func DeserializeUser(c *fiber.Ctx) error {
	var access_token string

	if c.Cookies("access_token") != "" {
		access_token = c.Cookies("access_token")
	}

	if access_token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "You are not logged in"})
	}

	//config, err := config.LoadConfig(".")

	UserID, err := database.RedisDb.Db.Get(access_token).Result()
	if err != nil {
		panic(err)
	}
	objID, err := primitive.ObjectIDFromHex(UserID)
	if err != nil {
		panic(err)
	}
	/*tokenClaims := handlers.ValidToken(access_token, userID)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	ctx := context.TODO()
	userid, err := database.RedisClient.Get(ctx, tokenClaims.TokenUuid).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "Token is invalid or session has expired"})
	}

	var user models.User
	err = initializers.DB.First(&user, "id = ?", userid).Error

	if err == gorm.ErrRecordNotFound {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "the user belonging to this token no logger exists"})
	}

	c.Locals("user", models.FilterUserRecord(&user))
	c.Locals("access_token_uuid", tokenClaims.TokenUuid)

	return c.Next()
	*/
	return c.Status(200).JSON(fiber.Map{
		"foo: ": objID,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})

	} else {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
	}
}

package handlers

import (
	"context"
	"net/mail"
	"time"

	"github.com/aminkhn/mongo-rest-api/config"
	"github.com/aminkhn/mongo-rest-api/database"
	"github.com/aminkhn/mongo-rest-api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func getUserByEmail(e string) (*models.User, error) {
	db := database.GetDBCollection("users")
	var user models.User

	filter := bson.D{{Key: "email", Value: e}}

	if err := db.FindOne(context.TODO(), filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func getUserByUsername(u string) (*models.User, error) {
	db := database.GetDBCollection("users")
	var user models.User

	filter := bson.D{{Key: "username", Value: u}}

	if err := db.FindOne(context.TODO(), filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func isEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// Login gets user and password
func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Identity string `json:"identity"`
		Password string `json:"password"`
	}
	type UserData struct {
		ID       primitive.ObjectID `json:"id"`
		Username string             `json:"username"`
		Email    string             `json:"email"`
		Password string             `json:"password"`
	}
	input := new(LoginInput)
	var userData UserData

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	}

	identity := input.Identity
	pass := input.Password
	userModel, err := new(models.User), *new(error)

	if isEmail(identity) {
		userModel, err = getUserByEmail(identity)
	} else {
		userModel, err = getUserByUsername(identity)
	}

	if userModel == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "User not found", "data": err})
	} else {
		userData = UserData{
			ID:       userModel.ID,
			Username: userModel.Username,
			Email:    userModel.Email,
			Password: userModel.Password,
		}
	}

	if !CheckPasswordHash(pass, userData.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = userData.Username
	claims["user_id"] = userData.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
}

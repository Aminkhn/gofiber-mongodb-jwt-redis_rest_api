package handlers

import (
	"strconv"

	"github.com/aminkhn/mongo-rest-api/database"
	"github.com/aminkhn/mongo-rest-api/models"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ValidToken(t *jwt.Token, id string) bool {
	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))

	return uid == n
}

func ValidUser(id string, p string) bool {
	db := database.GetDBCollection("users")
	filter := bson.D{{Key: "_id", Value: id}}

	var result models.User
	db.FindOne(nil, filter).Decode(&result)

	if result.Username == "" {
		return false
	}
	if !CheckPasswordHash(p, result.Password) {
		return false
	}
	return true
}

func GetAllUser(c *fiber.Ctx) error {
	query := bson.M{}
	db := database.GetDBCollection("users")

	// get all records as a cursor
	cursor, err := db.Find(c.Context(), query)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	users := make([]models.User, 0)

	// iterate the cursor and decode each item into a User
	for cursor.Next(c.Context()) {
		user := models.User{}
		err := cursor.Decode(&user)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		users = append(users, user)
	}
	// return users list in JSON format
	return c.Status(200).JSON(fiber.Map{
		"data": users,
	})
}

// GetUser get a user
func GetUser(c *fiber.Ctx) error {
	// get id by params
	params := c.Params("id")

	_id, err := primitive.ObjectIDFromHex(params)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	filter := bson.D{{Key: "_id", Value: _id}}

	var result models.User

	if err := database.GetDBCollection("users").FindOne(c.Context(), filter).Decode(&result); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": result,
	})
}

// CreateUser new user
func CreateUser(c *fiber.Ctx) error {
	type NewUser struct {
		Name     string `json:"name"`
		Family   string `json:"family"`
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	collection := database.GetDBCollection("users")

	// New User struct
	user := new(models.User)
	// Parse body into struct
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})
	}

	// force MongoDB to always set its own generated ObjectIDs
	user.ID = primitive.NewObjectID()
	// hash password for security
	hash, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})
	}

	user.Password = hash
	// insert the record
	insertionResult, err := collection.InsertOne(c.Context(), user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// get the just inserted record in order to return it as response
	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), filter)

	// decode the Mongo record into user
	createdUsers := &models.User{}
	createdRecord.Decode(createdUsers)

	// return the created user in JSON format
	return c.Status(201).JSON(fiber.Map{
		"message": "success",
		"data":    createdUsers,
	})
}

// upsateUser update user put
func UpdateUserPut(c *fiber.Ctx) error {
	idParam := c.Params("id")
	userID, err := primitive.ObjectIDFromHex(idParam)

	// the provided ID might be invalid ObjectID
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user := new(models.User)
	// Parse body into struct
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Find the user and update its data
	query := bson.D{{Key: "_id", Value: userID}}
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "name", Value: user.Name},
				{Key: "family", Value: user.Family},
				{Key: "username", Value: user.Username},
				{Key: "email", Value: user.Email},
			},
		},
	}
	err = database.GetDBCollection("users").FindOneAndUpdate(c.Context(), query, update).Err()

	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return c.Status(404).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// return the updated user
	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

// UpdateUser update user patch
func UpdateUserPatch(c *fiber.Ctx) error {
	idParam := c.Params("id")
	userID, err := primitive.ObjectIDFromHex(idParam)

	// the provided ID might be invalid ObjectID
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user := new(models.User)
	// Parse body into struct
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Find the user and update its data
	query := bson.D{{Key: "_id", Value: userID}}
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "name", Value: user.Name},
				{Key: "family", Value: user.Family},
				{Key: "username", Value: user.Username},
				{Key: "email", Value: user.Email},
			},
		},
	}
	err = database.GetDBCollection("users").FindOneAndUpdate(c.Context(), query, update).Err()

	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return c.Status(404).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// return the updated user
	return c.Status(200).JSON(fiber.Map{
		"message": "success",
	})
}

// DeleteUser delete user
func DeleteUser(c *fiber.Ctx) error {
	userID, err := primitive.ObjectIDFromHex(
		c.Params("id"),
	)

	// the provided ID might be invalid ObjectID
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// find and delete the employee with the given ID
	query := bson.D{{Key: "_id", Value: userID}}
	result, err := database.GetDBCollection("users").DeleteOne(c.Context(), &query)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// the user might not exist
	if result.DeletedCount < 1 {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// the record was deleted
	return c.Status(204).JSON(fiber.Map{
		"message": "Record succesfuly deleted from database!",
	})
}

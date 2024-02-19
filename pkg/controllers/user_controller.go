package controllers

import (
	"context"
	"fmt"
	"math"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/soulinmaikadua/go-with-mongodb/pkg/configs"
	"github.com/soulinmaikadua/go-with-mongodb/pkg/models"
	"github.com/soulinmaikadua/go-with-mongodb/pkg/utils"
)

type Users struct {
	Page     int                `json:"page"`
	Total    int                `json:"total"`
	UserList []models.UserModel `json:"data"`
}

func GetUsers(c *fiber.Ctx) error {
	limit, skip := utils.ParsePaginationParams(c)
	// get connection
	db := configs.GetConnect()

	totalCount, err := db.Collection("traders").CountDocuments(context.Background(), bson.D{})
	if err != nil {
		fmt.Println("Error getting")
		return c.Status(500).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	// get users
	query := bson.D{{}}
	cursor, err := db.Collection("traders").Find(context.Background(), query, options.Find().SetSkip(int64(skip)).SetLimit(int64(limit)))
	if err != nil {
		fmt.Println("can't find users")
		return c.Status(500).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	users := Users{
		Page:     0,
		Total:    0,
		UserList: []models.UserModel{},
	}

	// iterate the cursor and decode each item into an Users
	if err := cursor.All(context.Background(), &users.UserList); err != nil {
		fmt.Println("Decode failed")
		return c.Status(500).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	// Calculate the total number of users
	users.Total = int(totalCount)
	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))
	users.Page = totalPages

	return c.JSON(users)
}
func GetUser(c *fiber.Ctx) error {

	userID := c.Params("id")

	// get connection
	db := configs.GetConnect()
	query := bson.M{"_id": userID}

	fmt.Println(query)

	var user models.UserModel
	err := db.Collection("traders").FindOne(context.Background(), query).Decode(&user)
	if err != nil {
		// Check if the error indicates "document not found"
		if err == mongo.ErrNoDocuments {
			// Return a 404 Not Found response
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   true,
				"message": "User not found",
			})
		}
		fmt.Println("can't find users")
		return c.Status(500).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.JSON(user)
}

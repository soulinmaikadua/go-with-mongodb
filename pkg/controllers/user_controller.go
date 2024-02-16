package controllers

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/soulinmaikadua/go-with-mongodb/pkg/configs"
	"github.com/soulinmaikadua/go-with-mongodb/pkg/models"
)

func GetUsers(c *fiber.Ctx) error {
	// Connect to the database
	if err := configs.Connect(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("test users")
	db := configs.Mg.Db
	// get all records as a cursor
	query := bson.D{{}}
	cursor, err := db.Collection("traders").Find(c.Context(), query)
	if err != nil {
		fmt.Println("can't find users")
		return c.Status(500).SendString(err.Error())
	}

	var users []models.UserModel = make([]models.UserModel, 0)

	// iterate the cursor and decode each item into an Employee
	if err := cursor.All(c.Context(), &users); err != nil {
		fmt.Println("Decode failed")
		return c.Status(500).SendString(err.Error())

	}
	// return users list in JSON format
	return c.JSON(users)
}

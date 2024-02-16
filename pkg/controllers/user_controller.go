package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUsers(c *fiber.Ctx) error {
	// get all records as a cursor
	query := bson.D{{}}
	cursor, err := Mg.Db.Collection("traders").Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var users []UserModel = make([]UserModel, 0)

	// iterate the cursor and decode each item into an Employee
	if err := cursor.All(c.Context(), &users); err != nil {
		return c.Status(500).SendString(err.Error())

	}
	// return users list in JSON format
	return c.JSON(users)
}

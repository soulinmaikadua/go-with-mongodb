package controllers

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/soulinmaikadua/go-with-mongodb/pkg/configs"
	"github.com/soulinmaikadua/go-with-mongodb/pkg/models"
	"github.com/soulinmaikadua/go-with-mongodb/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func Signup(c *fiber.Ctx) error {
	user := &models.InputUser{}

	// Parse body into struct
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	// get connection
	db := configs.GetConnect()

	result, err := db.Collection("traders").InsertOne(context.Background(), user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	fmt.Println(result.InsertedID)
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {

	user := &models.InputUser{}
	inputUser := &models.LoginInput{}

	// Parse body into struct
	if err := c.BodyParser(inputUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	// get connection
	db := configs.GetConnect()

	err := db.Collection("traders").FindOne(context.Background(), bson.M{"email": inputUser.Email}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	// Verified Password

	match := utils.VerifyPassword(inputUser.Password, user.Password)

	if !match {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	user.Password = "FUCK_YOU"

	token, err := utils.GenerateNewToken(*user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	c.Set("token", token)
	return c.JSON(user)
}

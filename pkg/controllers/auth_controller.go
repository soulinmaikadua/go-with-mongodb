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

type LoginInput struct {
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

func Login(c *fiber.Ctx) error {

	user := &models.InputUser{}
	inputUser := &LoginInput{}

	// Check, if requested JSON is valid
	if err := c.BodyParser(inputUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	// get connection
	db := configs.GetConnect()

	err := db.Collection("traders").FindOne(context.Background(), bson.M{"email": inputUser.Email}).Decode(&user)
	// Check, if requested JSON is valid
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
	fmt.Println("VerifyPassword", match)
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

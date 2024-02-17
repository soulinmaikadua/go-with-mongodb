package utils

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ParsePaginationParams(c *fiber.Ctx) (int, int) {
	limitStr := c.Query("limit")
	pageStr := c.Query("page")

	// Parse the limit parameter to an integer, or set default value to 10
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10 // Set default value to 10 if limit is not provided or invalid
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1 // Default page number
	}

	// Calculate the skip value
	skip := (page - 1) * limit

	return limit, skip
}

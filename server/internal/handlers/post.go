package handlers

import (
	"github.com/gofiber/fiber/v2"

	"doodlegram/internal/data"
)

func GetPostsHandler(r data.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		posts, err := r.GetPosts()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error fetching books"})
		}
		return c.Status(fiber.StatusOK).JSON(posts)
	}
}

package handlers

import (
	"github.com/gofiber/fiber/v2"

	"doodlegram/internal/data"
)

func GetPostsHandler(r data.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		posts, err := r.GetPosts()
		if err != nil {
			return fiber.ErrInternalServerError
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"posts": posts})
	}
}

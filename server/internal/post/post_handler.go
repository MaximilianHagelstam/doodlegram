package post

import (
	"github.com/gofiber/fiber/v2"
)

func GetPostsHandler(service Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		posts, err := service.GetPosts()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error fetching books"})
		}
		return c.Status(fiber.StatusOK).JSON(posts)
	}
}

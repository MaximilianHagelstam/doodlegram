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

func GetPostByIDHandler(r data.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		post, err := r.GetPostByID(id)
		if err != nil {
			return fiber.ErrNotFound
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"post": post})
	}
}

func DeletePostHandler(r data.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		err := r.DeletePost(id)
		if err != nil {
			return fiber.ErrNotFound
		}
		return c.SendStatus(fiber.StatusNoContent)
	}
}

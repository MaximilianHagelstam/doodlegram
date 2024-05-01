package post

import (
	"github.com/gofiber/fiber/v2"
)

func PostRouter(app fiber.Router, service Service) {
	app.Get("/posts", GetPostsHandler(service))
}

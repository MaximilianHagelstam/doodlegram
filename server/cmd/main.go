package main

import (
	"doodlegram/internal/post"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	postRepository := post.NewRepo()
	postService := post.NewService(postRepository)

	api := app.Group("/api")
	post.PostRouter(api, postService)

	log.Fatal(app.Listen(":8080"))
}

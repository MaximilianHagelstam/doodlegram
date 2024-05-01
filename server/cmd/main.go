package main

import (
	"doodlegram/internal/database"
	"doodlegram/internal/post"
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	db := database.New()
	postRepository := post.NewRepo(db)
	postService := post.NewService(postRepository)

	api := app.Group("/api")
	post.PostRouter(api, postService)

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	err := app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}

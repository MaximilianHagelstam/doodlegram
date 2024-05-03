package main

import (
	"doodlegram/internal/data"
	"doodlegram/internal/handlers"
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

	db := data.NewDB()
	repo := data.NewRepo(db)

	api := app.Group("/api")

	api.Get("/posts", handlers.GetPostsHandler(repo))
	api.Get("/posts/:id", handlers.GetPostByIDHandler(repo))
	api.Delete("/posts/:id", handlers.DeletePostHandler(repo))

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	err := app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}

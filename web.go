package main

import (
	"dylangluck/blog"
	"dylangluck/handlers"
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

//go:embed views/*
var viewsfs embed.FS

func main() {
	// Posts Collection
	c := blog.New()

	// Create a new engine
	engine := html.NewFileSystem(http.FS(viewsfs), ".html")
	engine.AddFunc(
		// Add unescape function
		"unescape", func(s string) template.HTML {
			return template.HTML(s)
		},
	)

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Add middleware
	app.Use(cache.New())
	app.Use(recover.New())

	// Route: Index
	app.Get("/", handlers.HandleIndex(c))

	// Route: Blog
	app.Get("/blog", handlers.HandleBlog(c))
	app.Get("/blog/:name", handlers.HandlePost(c))
	app.Get("/blog/tag/:tag", handlers.HandleTag(c))

	// Serve static files with cache
	app.Static("/", "./public", fiber.Static{
		MaxAge: 3600,
	})

	// Handle not found
	app.Use(handlers.NotFound)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(app.Listen(":" + port))
}

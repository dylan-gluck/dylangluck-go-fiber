package main

import (
	"dylangluck/blog"
	"embed"
	"html/template"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

//go:embed views/*
var viewsfs embed.FS

func main() {
	// Create a new engine
	engine := html.NewFileSystem(http.FS(viewsfs), ".html")
	engine.AddFunc(
		// add unescape function
		"unescape", func(s string) template.HTML {
			return template.HTML(s)
		},
	)

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Serve static files
	app.Static("/", "./public")

	// Route: Index
	app.Get("/", func(c *fiber.Ctx) error {
		blog.GetPosts()
		// Render index within layouts/main
		return c.Render("views/index", fiber.Map{
			"Title": "Profile",
		}, "views/layouts/main")
	})

	app.Get("/blog/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		meta := blog.GetPostMeta(name)
		content := blog.GetPostHTML(name)
		// Render article within layouts/main
		return c.Render("views/post", fiber.Map{
			"Title":   meta.Title,
			"Short":   meta.Short,
			"Date":    meta.Date,
			"Tags":    meta.Tags,
			"Content": content,
		}, "views/layouts/main")
	})

	log.Fatal(app.Listen(":8080"))
}

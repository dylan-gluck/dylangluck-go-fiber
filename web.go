package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
    // Create a new engine
    engine := html.New("./views", ".html")

    // Pass the engine to the Views
    app := fiber.New(fiber.Config{
        Views: engine,
    })

    app.Static("/", "./public")

    app.Get("/", func(c *fiber.Ctx) error {
        // Render index within layouts/main
        return c.Render("index", fiber.Map{
						"Title": "Profile",
            "H1": "Dylan Navajas Gluck",
        }, "layouts/main")
    })

    log.Fatal(app.Listen(":8080"))
}

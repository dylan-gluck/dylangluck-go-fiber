package handlers

import (
	"dylangluck/blog"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Handler func(c *fiber.Ctx) error

func NotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).SendFile("public/404.html")
}

func HandleIndex(collection *blog.Posts) Handler {
	return func(c *fiber.Ctx) error {
		// Get recent posts
		posts, err := collection.Recent(2, 1)
		if err != nil {
			log.Default().Println(err)
		}

		// Render index within layouts/main
		return c.Render("views/index", fiber.Map{
			"Title": "Profile",
			"Posts": posts,
		}, "views/layouts/main")
	}
}

func HandleBlog(collection *blog.Posts) Handler {
	return func(c *fiber.Ctx) error {
		// Get recent posts
		// TODO: Pagination
		posts, err := collection.Recent(10, 1)
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendFile("public/404.html")
		}

		// Render blog within layouts/main
		return c.Render("views/blog", fiber.Map{
			"Title": "Blog",
			"Posts": posts,
		}, "views/layouts/main")
	}
}

func HandleTag(collection *blog.Posts) Handler {
	return func(c *fiber.Ctx) error {
		tag := c.Params("tag")

		filtered := collection.FilterByTag(tag)
		posts, err := filtered.Recent(10, 1)
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendFile("public/404.html")
		}

		// Render blog within layouts/main
		return c.Render("views/blog", fiber.Map{
			"Title": tag,
			"Posts": posts,
		}, "views/layouts/main")
	}
}

func HandlePost(collection *blog.Posts) Handler {
	return func(c *fiber.Ctx) error {
		name := c.Params("name")

		meta, err := collection.PostMeta(name)
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendFile("public/404.html")
		}
		content, err := collection.PostContent(name)
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendFile("public/404.html")
		}

		// Render article within layouts/main
		return c.Render("views/post", fiber.Map{
			"Id":      meta.Id,
			"Title":   meta.Title,
			"Short":   meta.Short,
			"Date":    meta.Date,
			"Tags":    meta.Tags,
			"Content": content,
		}, "views/layouts/main")
	}
}

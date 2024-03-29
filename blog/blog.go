package blog

import (
	"dylangluck/util"
	"log"
	"os"
	"slices"
	"time"

	yaml "github.com/goccy/go-yaml"
	"github.com/gofiber/fiber/v2"
)

type Collection interface {
	ReadDir()
	SortByDate()
	Recent(limit int, page int) (Posts, error)
	PostMeta(name string) (Post, error)
	PostContent(name string) (string, error)
}

type Post struct {
	Id    string
	Title string
	Short string
	Date  string
	Tags  []string
}

type Posts []Post

// Collection of posts
func New() *Posts {
	p := &Posts{}
	p.ReadDir()
	p.SortByDate()
	return p
}

func (p *Posts) ReadDir() {
	files, err := os.ReadDir("blog/articles/")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		post := getPostMeta(file.Name())
		*p = append(*p, post)
	}
}

func (p *Posts) SortByDate() {
	for i := 0; i < len(*p); i++ {
		for j := i + 1; j < len(*p); j++ {
			if parseDate((*p)[i].Date).Before(parseDate((*p)[j].Date)) {
				(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
			}
		}
	}
}

func (p *Posts) Recent(limit int, page int) (Posts, error) {
	start := limit * (page - 1)
	if start > len(*p) {
		return Posts{}, fiber.ErrNotFound
	}

	end := limit * page
	if end > len(*p) {
		end = len(*p)
	}

	return (*p)[start:end], nil
}

func (p *Posts) PostMeta(name string) (Post, error) {
	id := slices.IndexFunc(*p, func(post Post) bool {
		return post.Id == name
	})
	if id == -1 {
		return Post{}, fiber.ErrNotFound
	}
	return (*p)[id], nil
}

func (p *Posts) PostContent(name string) (string, error) {
	md, err := os.ReadFile("blog/articles/" + name + "/post.md")
	if err != nil {
		return "", fiber.ErrNotFound
	}
	return string(util.MdToHTML([]byte(md))), nil
}

// Get post meta
func getPostMeta(name string) Post {
	yml, err := os.ReadFile("blog/articles/" + name + "/meta.yml")
	if err != nil {
		log.Panic(err)
	}

	var post Post
	post.Id = name
	if err := yaml.Unmarshal([]byte(yml), &post); err != nil {
		log.Panic(err)
	}
	return post
}

// Convert date string to time.Time
func parseDate(date string) time.Time {
	const format = "2006-01-02"
	t, err := time.Parse(format, date)
	if err != nil {
		log.Panic(err)
	}
	return t
}

// Sort posts by .Date
func sortPostsByDate(posts Posts) Posts {
	for i := 0; i < len(posts); i++ {
		for j := i + 1; j < len(posts); j++ {
			if parseDate(posts[i].Date).Before(parseDate(posts[j].Date)) {
				posts[i], posts[j] = posts[j], posts[i]
			}
		}
	}
	return posts
}

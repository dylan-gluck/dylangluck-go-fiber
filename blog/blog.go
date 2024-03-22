package blog

import (
	"dylangluck/util"
	"log"
	"os"
	"time"

	yaml "github.com/goccy/go-yaml"
)

type Post struct {
	Id    string
	Title string
	Short string
	Date  string
	Tags  []string
}

type Posts []Post

// Get post meta
func GetPostMeta(name string) Post {
	yml, err := os.ReadFile("blog/articles/" + name + "/meta.yml")
	if err != nil {
		log.Fatal(err)
	}

	var post Post
	post.Id = name
	if err := yaml.Unmarshal([]byte(yml), &post); err != nil {
		log.Fatal(err)
	}
	return post
}

// Get all posts
func GetPosts() Posts {
	files, err := os.ReadDir("blog/articles/")
	if err != nil {
		log.Fatal(err)
	}

	var posts Posts
	for _, file := range files {
		post := GetPostMeta(file.Name())
		posts = append(posts, post)
	}
	return posts
}

// Get recent posts
func GetRecentPosts(limit int) Posts {
	posts := GetPosts()
	sorted := sortPostsByDate(posts)
	if len(sorted) > limit {
		return sorted[:limit]
	}
	return sorted
}

// Convert date string to time.Time
func parseDate(date string) time.Time {
	const format = "2006-01-02"
	t, err := time.Parse(format, date)
	if err != nil {
		log.Fatal(err)
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

// Render single post
func RenderPostContent(name string) string {
	md, err := os.ReadFile("blog/articles/" + name + "/post.md")
	if err != nil {
		log.Fatal(err)
	}
	return string(util.MdToHTML([]byte(md)))
}

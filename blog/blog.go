package blog

import (
	"dylangluck/util"
	"log"
	"os"

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

func GetPostMeta(name string) Post {
	yml, err := os.ReadFile("blog/articles/" + name + "/meta.yml")
	if err != nil {
		log.Fatal(err)
	}

	var post Post
	if err := yaml.Unmarshal([]byte(yml), &post); err != nil {
		log.Fatal(err)
	}
	return post
}

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

func GetPostHTML(name string) string {
	md, err := os.ReadFile("blog/articles/" + name + "/post.md")
	if err != nil {
		log.Fatal(err)
	}
	return string(util.MdToHTML([]byte(md)))
}

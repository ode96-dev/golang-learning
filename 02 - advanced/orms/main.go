package main

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

/*
TODO: ORMs in Go
- simplify database operations by letting us work with Go structs instead of raw SQL.
	- making code cleaner, type-safe, and easier to maintain.
- popular ORMS in Go - GORM (widely adopted), Ent, SQLBoiler, XORM
*/

type Post struct {
	gorm.Model
	Title string
	Slug  string `gorm:"uniqueIndex:idx_slug"`
	Likes uint
}

func (p Post) String() string {
	return fmt.Sprintf("Post Title: %s, Slug: %s", p.Title, p.Slug)
}

var db, _ = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

func main() {
	// db.AutoMigrate(&Post{})
	// db.Migrator().AddColumn(&Post{}, "likes")
	// freshPost := createPost("new slug", "new-slug-1")

	// fmt.Println(freshPost)

	oldPost := getPost("new-slug-1")
	fmt.Println(oldPost)
}

func createPost(title string, slug string) Post {
	newPost := Post{Title: title, Slug: slug}

	if res := db.Create(&newPost); res.Error != nil {
		panic(res.Error)
	}

	return newPost
}

func getPost(slug string) Post {
	targetPost := Post{Slug: slug}

	if res := db.First(&targetPost); res.Error != nil {
		panic(res.Error)
	}

	return targetPost

}

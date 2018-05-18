package main

import (
	"fmt"
)

/*type ResultPost struct{
	Content string
	Author string
}*/

func SelectPostByID(id int) {
	fmt.Println("Select Post By  PostID")
	var posts []Post
	err := Db.Model(&posts).
		Column("post.*", "Author").
		Where("post.id = ?", id).
		Select()
	if err != nil {
		panic(err)
	}

	for _, post := range posts {
		fmt.Println(post.Id, "|", post.Author.Name, "|", post.Content)
	}
}

func SelectPostByAuthor(authorid int) {
	fmt.Println("Select Posts By AuthorID")
	var posts []Post;
	err := Db.Model(&posts).
		Column("post.*", "Author").
		Where("author.id = ?", authorid).
		Select()
	if err != nil {
		panic(err)
	}
	for _, post := range posts {
		fmt.Println(post.Id, "|", post.Author.Name, "|", post.Content)
	}
}
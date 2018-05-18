package main

import (
	"fmt"
)

/*type ResultPost struct{
	Content string
	Author string
}*/

func SelelectPostByID(id int) {
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
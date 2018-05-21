package main

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"fmt"
)

type Post struct {
	Id int //Primary key
	Content string
	AuthorID int  //Foreign key trỏ từ bảng User.Id
	Author *User  //Author
}

type User struct {
	Id int
	Name string
}

var Db *pg.DB

func ConnectDB() error {
	Db = pg.Connect(&pg.Options{
		User: "postgres",
		Password: "123",
		Database: "demo",
		Addr: "localhost:5432",
	})

	var n int
	_, err := Db.QueryOne(pg.Scan(&n), "SELECT 1")
	return err
}

func InitSchema() {
	var post Post
	var user User

	for _, model := range []interface{} {&user, &post} {

		Db.DropTable(model, &orm.DropTableOptions{
			IfExists:true,
			Cascade:true,
		})

		err := Db.CreateTable(model, &orm.CreateTableOptions{
			Temp: false,
			FKConstraints: true,
			IfNotExists: true,
		})

		if err != nil {
			panic(err)
		}
	}
}

func SaveData() {
	user1 := User{
		Name: "Nguyễn Hàn Duy",
	}

	user2 := User {
		Name: "Nguyễn Thành Long",
	}

	user3 := User {
		Name: "Nguyễn Thành Luỹ",
	}

	err := Db.Insert(&user1, &user2, &user3)
	if err != nil {
		fmt.Println("err")
	}

	fmt.Println(user1.Id)
	fmt.Println(user2.Id)
	fmt.Println(user3.Id)

	post1 := Post {
		Content: "Anh yêu em",
		AuthorID: user1.Id,
	}

	post2 := Post {
		Content: "Anh chán em",
		AuthorID: user2.Id,
	}

	post3 := Post {
		Content: "Anh thích em",
		AuthorID: user3.Id,
	}

	post4 := Post {
		Content: "Anh cực thích em",
		AuthorID: user1.Id,
	}

	err = Db.Insert(&post1, &post2, &post3, &post4)
	if err != nil {
		fmt.Println("err")
	}
	fmt.Println(post1.Author)
}
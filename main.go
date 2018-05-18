package main

import (
	"fmt"
)

func main() {
	fmt.Println("Demo Go-PG")
	if err := ConnectDB(); err != nil {
		fmt.Println(err)
	}

	InitSchema()
	SaveData()

	SelectPostByID(1)
	SelectPostByAuthor(1)
	defer Db.Close()
}

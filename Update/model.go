package main

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type Student struct {
	Id int //Primary key
	Name string
	Age int
}

var Db *pg.DB

func ConnectDB() error {
	Db = pg.Connect(&pg.Options{
		User: "postgres",
		Password: "123",
		Database: "Update",
		Addr: "localhost:5432",
	})

	var n int
	_, err := Db.QueryOne(pg.Scan(&n), "SELECT 1")
	return err
}

func InitSchema() {
	var student Student

	for _, model := range []interface{}{&student} {
		Db.DropTable(model, &orm.DropTableOptions{
			IfExists:true,
			Cascade:true,
		})
		err := Db.CreateTable(model, &orm.CreateTableOptions{
			Temp: false,
			IfNotExists: true,
		})

		if err != nil {
			panic(err)
		}
	}
}

func SaveData() {
	queryString := []string{
		"INSERT INTO students VALUES (1, 'Duy'), (2, 'Long'), (3, 'Huy'), (4, 'Luy'), (5, 'Thanh'), (6, 'Duong'), (7, 'Quyen')",
	}
	for _, query := range queryString {
		_, err := Db.Exec(query)
		if err != nil {
			panic(err)
		}
	}
}
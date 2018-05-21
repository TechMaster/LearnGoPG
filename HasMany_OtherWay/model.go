package main

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type Phone struct {
	Id int // Primary key
	Number string
	StudentId int
}

type Student struct {
	Id int //Primary key
	Name string
	Phones []Phone
}

var Db *pg.DB

func ConnectDB() error {
	Db = pg.Connect(&pg.Options{
		User: "postgres",
		Password: "123",
		Database: "hello",
		Addr: "localhost:5432",
	})

	var n int
	_, err := Db.QueryOne(pg.Scan(&n), "SELECT 1")
	return err
}

func InitSchema() {
	var student Student
	var phone Phone

	for _, model := range []interface{}{&phone, &student} {
		Db.DropTable(model, &orm.DropTableOptions{
			IfExists:true,
			Cascade:true,
		})
		err := Db.CreateTable(model, &orm.CreateTableOptions{
			Temp: false,
			IfNotExists: true,
			FKConstraints: true,
		})

		if err != nil {
			panic(err)
		}
	}
}

func SaveData() {
	queryString := []string{
		"INSERT INTO phones VALUES (1, '123', 1), (2, '456', 1), (3, '789', 2), (4, '56789', 2), (5, '5432', 3)",
		"INSERT INTO students VALUES (1, 'Duy'), (2, 'Long'), (3, 'Huy')",
	}
	for _, query := range queryString {
		_, err := Db.Exec(query)
		if err != nil {
			panic(err)
		}
	}
}
package main

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type Student struct {
	Id int //Primary key
	Name string
	AccountId int
	Account *Account
}

type Account struct {
	Id int // Primary key
	Number string
}

var Db *pg.DB

func ConnectDB() error {
	Db = pg.Connect(&pg.Options{
		User: "postgres",
		Password: "123",
		Database: "OneToOne",
		Addr: "localhost:5432",
	})

	var n int
	_, err := Db.QueryOne(pg.Scan(&n), "SELECT 1")
	return err
}

func InitSchema() {
	var student Student
	var account Account

	for _, model := range []interface{}{&account, &student} {
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
		"INSERT INTO accounts VALUES (1, 'Duy123'), (2, 'Long456'), (3, 'Huy789')",
		"INSERT INTO students VALUES (1, 'Duyhaha', 1), (2, 'Longhehe', 1), (3, 'Huyhoho', 3)",
	}
	for _, query := range queryString {
		_, err := Db.Exec(query)
		if err != nil {
			panic(err)
		}
	}
}
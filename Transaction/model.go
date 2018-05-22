package main

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type Transaction struct {
	Counter int
}

var Db *pg.DB

func ConnectDB() error {
	Db = pg.Connect(&pg.Options{
		User: "postgres",
		Password: "123",
		Database: "Transaction",
		Addr: "localhost:5432",
	})

	var n int
	_, err := Db.QueryOne(pg.Scan(&n), "SELECT 1")
	return err
}

func InitSchema() {
	var transaction Transaction

	for _, model := range []interface{}{&transaction} {
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
		"INSERT INTO transactions VALUES (0)",
	}
	for _, query := range queryString {
		_, err := Db.Exec(query)
		if err != nil {
			panic(err)
		}
	}
}
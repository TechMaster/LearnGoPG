package main

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type Person struct {
	Id int //Primary key
	Name string
}

type Phone struct {
	Id int // Primary key
	Number string
	PersonId int  // Foreign key tro den bang Person
	Person *Person
}

var Db *pg.DB

func ConnectDB() error {
	Db = pg.Connect(&pg.Options{
		User: "postgres",
		Password: "123",
		Database: "OneToMany",
		Addr: "localhost:5432",
	})

	var n int
	_, err := Db.QueryOne(pg.Scan(&n), "SELECT 1")
	return err
}

func InitSchema() {
	var person Person
	var phone Phone

	for _, model := range []interface{}{&phone, &person} {
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
	// Insert person
	person1 := Person{
		Id: 1,
		Name:"Duy",
	}
	person2 := Person{
		Id: 2,
		Name:"Long",
	}
	person3 := Person{
		Id: 3,
		Name:"Huy",
	}
	errInsertPerson := Db.Insert(&person1,&person2,&person3)
	if errInsertPerson != nil {
		panic(errInsertPerson)
	}

	// Insert phone
	phone1 := Phone{
		Id: 1,
		Number:"123",
		PersonId: person1.Id,
	}
	phone2 := Phone{
		Id: 2,
		Number:"456",
		PersonId: person1.Id,
	}
	phone3 := Phone{
		Id: 3,
		Number:"789",
		PersonId: person1.Id,
	}
	phone4 := Phone{
		Id: 4,
		Number:"09876",
		PersonId: person2.Id,
	}
	phone5 := Phone{
		Id: 5,
		Number:"012345",
		PersonId: person2.Id,
	}
	phone6 := Phone{
		Id: 6,
		Number:"54321",
		PersonId: person3.Id,
	}
	errInsertPhone := Db.Insert(&phone1,&phone2,&phone3, &phone4, &phone5, &phone6)
	if errInsertPhone != nil {
		panic(errInsertPhone)
	}
}
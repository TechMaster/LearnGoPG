package main

import (
	"log"
	// "github.com/go-pg/pg/orm"
)

func MappingPersonToPhone() {
	var students []Student
	err := Db.Model(&students).
		Column("student.*").
		Relation("Account").
		Select()
	if err != nil {
		panic(err)
	}

	for _, student := range students {
		log.Println("Student", student.Id, student.Name, "|", "Account", student.Account.Number)
	}
}
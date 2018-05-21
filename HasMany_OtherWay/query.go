package main

import (
	"log"
	// "github.com/go-pg/pg/orm"
)

func MappingPersonToPhone() {
	var students []Student
	err := Db.Model(&students).
		Column("student.*").
		Relation("Phones").
		Select()
	if err != nil {
		panic(err)
	}

	for _, student := range students {
		for _, phone := range student.Phones {
			log.Println("Student", student.Id, student.Name, "|", "Phone", phone.Number)
		}
	}
}
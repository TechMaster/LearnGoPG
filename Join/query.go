package main

import (
	"log"
	// "github.com/go-pg/pg/orm"
)

func MappingPersonToPhone() {
	var phones []Phone
	var result []struct {
		PhoneId int // Primary key
		PhoneNumber string
		StudentId int
		StudentName string
	}

	err := Db.Model(&phones).
	ColumnExpr("phone.id AS phone_id, phone.number AS phone_number, phone.student_id AS student_id").
	ColumnExpr("students.name AS student_name").
	Join("RIGHT JOIN students ON students.id = phone.student_id").
	Select(&result)
	
	if err != nil {
		panic(err)
	}

	log.Println("PhoneId", "PhoneNumber", "StudentId", "StudentName")

	for _, res := range result {
		// log.Printf("Student %d has %d phone number(s)\n", res.StudentID, res.PhoneCount)
		log.Println(res.PhoneId, "     ",  res.PhoneNumber, "          ", res.StudentId, "          ", res.StudentName)
	}
}
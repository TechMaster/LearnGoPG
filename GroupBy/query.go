package main

import (
	"log"
	// "github.com/go-pg/pg/orm"
)

func MappingPersonToPhone() {
	var phones []Phone
	var queryResults []struct {
		StudentID  int
		PhoneCount int
	}

	err := Db.Model(&phones).
		Column("student_id").
		ColumnExpr("count(number) AS phone_count").
		Group("student_id").
		Order("student_id").
		Select(&queryResults)
	if err != nil {
		panic(err)
	}

	for _, res := range queryResults {
		log.Printf("Student %d has %d phone number(s)\n", res.StudentID, res.PhoneCount)
	}
}
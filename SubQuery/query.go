package main

import (
	"log"
)

func SubQuery() {
	var phones []Phone
	var sum struct {
		Total int
	}

	studentPhone := Db.Model(&phones).
		Column("student_id").
		ColumnExpr("count(number) AS phone_count").
		Group("student_id").
		Order("student_id")

	err := Db.Model().With("student_phone", studentPhone).Table("student_phone").
		ColumnExpr("sum(student_phone.phone_count) AS total").Select(&sum)

	if err != nil {
		panic(err)
	}

	log.Println(sum.Total)
}
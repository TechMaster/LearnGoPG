package main

import (
	"log"
)

func MappingStudentToCourse() {
	// Mapping course to student
	var courses []Course
	var courseToStudent []struct {
		CourseId int
		CourseName string
		StudentId int
		StudentName string
	}
	err := Db.Model(&courses).
	ColumnExpr("course.id AS course_id, course.title AS course_name, course.student_id AS student_id").
	ColumnExpr("students.name AS student_name").
	Join("LEFT JOIN students ON students.id = course.student_id").
	Select(&courseToStudent)
	if err != nil {
		panic(err)
	}
	for _,res := range courseToStudent {
		log.Println(res)
	}

	log.Println("---------------------------------------------------------")

	// Mapping student to course
	var students []Student
	var studentToCourse []struct {
		StudentId int
		StudentName string
		CourseId int
		CourseName string
	}
	err1 := Db.Model(&students).
	ColumnExpr("student.id AS student_id, student.name AS student_name, student.course_id AS course_id").
	ColumnExpr("courses.title AS course_name").
	Join("LEFT JOIN courses ON student.course_id = courses.id").
	Select(&studentToCourse)
	if err1 != nil {
		panic(err)
	}

	for _,result := range studentToCourse {
		log.Println(result)
	}
}


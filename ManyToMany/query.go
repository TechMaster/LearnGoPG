package main

import (
	"log"
)

func MappingStudentToCourse() {
	var courses []Course
	err := Db.Model(&courses).Relation("Students").Select()
	if err != nil {
		panic(err)
	}
	for _,course := range courses {
		for _, student := range course.Students {
			log.Println("Course",course.Id, course.Title, course.Teacher, student.Id, student.Name, student.Age, student.City)
		}
		// log.Println("Course",course.Id, course.Title, course.Students)
	}

	log.Println("---------------------------------------------")

	var students []Student
	err1 := Db.Model(&students).Relation("Courses").Select()
	if err1 != nil {
		panic(err1)
	}
	for _,student := range students {
		for _, course := range student.Courses {
			log.Println("Student",student.Id, student.Name,  student.Age, student.City, course.Id, course.Title, course.Teacher)	
		}
		// log.Println("Student",student.Id, student.Name, student.Courses)
	}

	log.Println("---------------------------------------------")
}

func CountNumerOfCourse() {
	var students []Student
	err := Db.Model(&students).Relation("Courses").
		Column("id", 
			"age", 
		).Select()
	if err != nil {
		panic(err)
	}
	
	for _,student := range students {
		log.Println(student)
	}	
}


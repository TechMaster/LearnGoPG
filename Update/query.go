package main

import (
	"log"
)

func Update() {
	student := &Student{Id: 1}
	err := Db.Select(student)
	if err != nil {
		panic(err)
	}

	student.Name = "Duy dau troc"
	err = Db.Update(student)
	if err != nil {
		panic(err)
	}

	err = Db.Select(student)
	if err != nil {
		panic(err)
	}

	log.Println(student)
}

func MultiUpdate() {
	student2 := &Student{
		Id: 2,
		Age: 30,
	}
	student3 := &Student{
		Id: 3,
		Age: 40,
	}

	_, err := Db.Model(student2, student3).Column("age").Update()
	if err != nil {
		panic(err)
	}
}

func MultiUpdateSlice() {
	students := []Student{
		{
			Id: 1,
			Name: "Duong Ukraine",
			Age: 20,
		},
		{
			Id: 2,
			Name: "Quyen Luc",
			Age: 21,
		},
	}

	_, err := Db.Model(&students).Column("name", "age").Update()
	if err != nil {
		panic(err)
	}
}
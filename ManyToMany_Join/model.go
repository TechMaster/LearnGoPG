package main

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type Student struct {
	Id int //Primary key
	Name string
	Age int
	City string
	CourseId int
	Courses []Course
}

type Course struct {
	Id int // Primary key
	Title string
	Teacher string
	StudentId int
	Students []Student
}

var Db *pg.DB

func ConnectDB() error {
	Db = pg.Connect(&pg.Options{
		User: "postgres",
		Password: "123",
		Database: "many",
		Addr: "localhost:5432",
	})

	var n int
	_, err := Db.QueryOne(pg.Scan(&n), "SELECT 1")
	return err
}

func InitSchema() {
	var student Student
	var course Course

	for _, model := range []interface{}{&student, &course} {
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
		"INSERT INTO courses VALUES (1, 'Docker can ban', 'Cuong', 1), (2, 'Linux qua cac vi du', 'Huy', 1), (3, 'React Native', 'Hao', 2), (4, 'Java can ban', 'Hiep', 2), (5, 'PHP Laravel', 'Hung', 3)",
		"INSERT INTO students VALUES (1, 'Duyhaha', 25, 'Hanoi', 1), (2, 'Longhehe', 24, 'London', 1), (3, 'Huyhoho', 27, 'Paris', 2)",
	}
	for _, query := range queryString {
		_, err := Db.Exec(query)
		if err != nil {
			panic(err)
		}
	}
}
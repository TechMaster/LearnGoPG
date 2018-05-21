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
	Courses []Course `pg:"many2many:student_to_courses"`
}

type Course struct {
	Id int // Primary key
	Title string
	Teacher string
	Students []Student `pg:"many2many:student_to_courses"`
}

type StudentToCourse struct {
	StudentId int
	CourseId int
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
	var studentToCourse StudentToCourse

	for _, model := range []interface{}{&student, &course, &studentToCourse} {
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
		"INSERT INTO courses VALUES (1, 'Docker can ban', 'Cuong'), (2, 'Linux qua cac vi du', 'Huy'), (3, 'React Native', 'Hao')",
		"INSERT INTO students VALUES (1, 'Duyhaha', 25, 'Hanoi'), (2, 'Longhehe', 24, 'London'), (3, 'Huyhoho', 27, 'Paris')",
		"INSERT INTO student_to_courses VALUES (1, 1), (1, 2), (3, 2)",
	}
	for _, query := range queryString {
		_, err := Db.Exec(query)
		if err != nil {
			panic(err)
		}
	}
}
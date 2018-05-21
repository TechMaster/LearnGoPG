package main

import (
    "fmt"
    "github.com/go-pg/pg/orm"
)

func init() {
	// Register many to many model so ORM can better recognize m2m relation.
	// This should be done before dependant models are used.
	orm.RegisterTable((*StudentToCourse)(nil))
}

func main() {
    fmt.Println("Demo Go-PG")
    if err := ConnectDB(); err != nil {
        fmt.Println(err)
    }
    // Viet code o day nhe!

    InitSchema()

    SaveData()

    MappingStudentToCourse()

    // CountNumerOfCourse()

    defer Db.Close()
}
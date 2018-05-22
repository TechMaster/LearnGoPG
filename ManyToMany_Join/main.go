package main

import (
    "fmt"
)

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
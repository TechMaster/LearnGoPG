package main

import (
	"log"
)

func MappingPersonToPhone() {
	var phones []Phone
	err1 := Db.Model(&phones).Column("Person", "phone.*").Select()
	if err1 != nil {
		panic(err1)
	}
	for _,phone := range phones {
		log.Println("Person", phone.Person.Id, phone.Person.Name ,"|","Phone", phone.Id, phone.Number)
	}	
}


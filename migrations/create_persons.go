package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/tangosource/martini_example/models"
)

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := gorm.Open("postgres", "dbname=martini_example sslmode=disable")
	PanicIf(err)
	db.CreateTable(person.Person{})
}

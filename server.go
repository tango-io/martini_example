package main

import (
	_ "fmt"

	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/martini-contrib/render"
	"github.com/tangosource/martini_example/controllers"
)

func SetupDB() gorm.DB {
	db, err := gorm.Open("postgres", "dbname=martini_example sslmode=disable")
	PanicIf(err)
	return db
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	m := martini.Classic()
	m.Map(SetupDB())
	m.Use(render.Renderer(render.Options{
		Directory:  "views",
		Layout:     "layouts/layout",
		Charset:    "UTF-8",
		IndentJSON: true,
		IndentXML:  true,
	}))

	m.Get("/persons/new", persons.NewPerson)
	m.Get("/persons/:id/edit", persons.EditPerson)
	m.Get("/persons/:id", persons.ShowPerson)
	m.Post("/persons/:id", persons.UpdatePerson)
	m.Get("/persons/:id/delete", persons.DeletePerson)
	m.Post("/persons", persons.CreatePerson)
	m.Get("/", persons.Index)
	m.Run()
}

package main

import (
	_ "fmt"
	"net/http"

	"strconv"

	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/martini-contrib/render"
	"github.com/tangosource/martini_example/models"
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

func index(r render.Render, req *http.Request, db gorm.DB) {
	persons := []person.Person{}
	db.Find(&persons)

	r.HTML(200, "index", persons)
}

func newPerson(r render.Render) {
	r.HTML(200, "persons/new", nil)
}

func createPerson(r render.Render, req *http.Request, db gorm.DB) {
	p := person.Person{}
	db.Create(&p)
	r.Redirect("/")
}

func editPerson(r render.Render, params martini.Params, db gorm.DB) {
	p := person.Person{}
	db.First(&p, params["id"])
	r.HTML(200, "persons/edit", p)
}

func updatePerson(r render.Render, req *http.Request, params martini.Params, db gorm.DB) {
	p := person.Person{}
	db.First(&p, params["id"])
	p.Name = req.FormValue("person[name]")
	p.Age, _ = strconv.Atoi(req.FormValue("person[age]"))
	p.Email = req.FormValue("person[email]")
	p.Job = req.FormValue("person[job]")
	db.Save(&p)
	r.Redirect("/")
}

func deletePerson(r render.Render, req *http.Request, params martini.Params, db gorm.DB) {
	p := person.Person{}
	db.First(&p, params["id"])
	db.Delete(&p)
	r.Redirect("/")
}

func showPerson(r render.Render, params martini.Params, db gorm.DB) {
	p := person.Person{}
	db.First(&p, params["id"])
	r.HTML(200, "persons/show", p)
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

	m.Get("/persons/new", newPerson)
	m.Get("/persons/:id/edit", editPerson)
	m.Get("/persons/:id", showPerson)
	m.Post("/persons/:id", updatePerson)
	m.Get("/persons/:id/delete", deletePerson)
	m.Post("/persons", createPerson)
	m.Get("/", index)
	m.Run()
}

package persons

import (
	"net/http"
	"strconv"

	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/render"
	"github.com/tangosource/martini_example/models"
)

func Index(r render.Render, req *http.Request, db gorm.DB) {
	persons := []person.Person{}
	db.Find(&persons)

	r.HTML(200, "index", persons)
}

func NewPerson(r render.Render) {
	r.HTML(200, "persons/new", nil)
}

func CreatePerson(r render.Render, req *http.Request, db gorm.DB) {
	age, _ := strconv.Atoi(req.FormValue("person[age]"))
	p := person.Person{
		Name:  req.FormValue("person[name]"),
		Age:   age,
		Email: req.FormValue("person[email]"),
		Job:   req.FormValue("person[job]"),
	}
	db.Create(&p)
	r.Redirect("/")
}

func EditPerson(r render.Render, params martini.Params, db gorm.DB) {
	p := person.Person{}
	db.First(&p, params["id"])
	r.HTML(200, "persons/edit", p)
}

func UpdatePerson(r render.Render, req *http.Request, params martini.Params, db gorm.DB) {
	p := person.Person{}
	db.First(&p, params["id"])
	age, _ := strconv.Atoi(req.FormValue("person[age]"))

	db.Model(&p).Updates(
		person.Person{
			Name:  req.FormValue("person[name]"),
			Age:   age,
			Email: req.FormValue("person[email]"),
			Job:   req.FormValue("person[job]"),
		},
	)
	r.Redirect("/")
}

func DeletePerson(r render.Render, req *http.Request, params martini.Params, db gorm.DB) {
	p := person.Person{}
	db.First(&p, params["id"])
	db.Delete(&p)
	r.Redirect("/")
}

func ShowPerson(r render.Render, params martini.Params, db gorm.DB) {
	p := person.Person{}
	db.First(&p, params["id"])
	r.HTML(200, "persons/show", p)
}

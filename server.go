package main

import (
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
)

type Person struct {
  Name      string
  Age       int
  Emails    []string
  Jobs      []*Job
}

type Job struct {
  Employer string
  Role     string
}

func main() {
  m := martini.Classic()

  m.Use(render.Renderer(render.Options {
    Directory: "views",
    Layout: "layouts/layout",
    Charset: "UTF-8",
    IndentJSON: true,
    IndentXML: true,
  }))

  m.Get("/", func(r render.Render) {

    person := Person{
      Name:   "Antonio Chavez",
      Age:    24,
      Emails: []string{ "cavjzz@gmail.com", "antonio.chavez@tangosource.com", "antonio@queuetechnologies.com" },
      Jobs:   []*Job{ &Job{ "TangoSource LLC", "Technical Leader" }, &Job{ "Queue Technologies Inc.", "Lead Engineer" } },
    }
    r.HTML(200, "index", person)
  })

  m.Run()
}

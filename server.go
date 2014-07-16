package main

import (
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
)


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
    r.HTML(200, "index", "Antonio")
  })

  m.Run()
}

# Martini Example

This application is a simple CRUD application written in [Golang](http://golang.org/)
using [Martini](http://martini.codegangsta.io/) and [PostgreSQL](http://www.postgresql.org/).


## Setting up martini for the project

In order to setup martini on the project you need to add the package to the imports:

``` go
import (
  // ...
  "github.com/go-martini/martini"
  // ...
)
```

Then you need to start martini on your application's main function:

``` go
func main() {
  // ...
  m := martini.Classic()
  m.Run()
  // ...
}
```

And that's it you have your martini application running, just run:

``` bash
go run server.go
```

And you can open it in your browser at `http://localhost:3000/`

## Writing routes

Next we're going to start writing routes, for martini to write
the routes is necessary to use the martini instance previously
defined and use the HTTP verb we want as following:

``` go
func main() {
  m := martini.Classic()
  m.Get("/", func(r render.Render) {
    return "Hello world!"
  })
  m.Run()
}
```

And you'll see the "Hello world!" in your browser
if you visit the index path, this same structure
works for the verbs "Put", "Post", "Delete".

## Rendering views

## Reading data from a postgresql database

## Inserting data to a postgresql database

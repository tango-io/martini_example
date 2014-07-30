# Martini Example

This application is a simple CRUD application written in [Golang](http://golang.org/)
using [Martini](http://martini.codegangsta.io/) and [PostgreSQL](http://www.postgresql.org/).

## Instalation

You can download Go from the [source](http://golang.org/doc/install)
or if you prefer you can use this [go version manager](https://github.com/moovweb/gvm)

Using gvm

``` bash
gvm install go1
```

The gopath will be automatically setup for your project (including vim files).
Next you need to install the packages:

``` bash
go get github.com/go-martini/martini
go get github.com/martini-contrib/render
go get github.com/lib/pq
go get github.com/codegangsta/gin
```

And run the server:

``` bash
go run server.go
# or
gin # This one has live reload feature
```


# Tutorial

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

In order to render views, you need to import the package to render first:

``` go
import (
  // ...
    "github.com/martini-contrib/render"
  // ...
)
```
By default martini's render look for your templates/, but you can change
the configuration as your convinience

``` go
func main() {
  m := martini.Classic()
  m.Use(render.Renderer(render.Options {
    Directory: "views",
    Layout: "layouts/layout",
    Charset: "UTF-8",
    IndentJSON: true,
    IndentXML: true,
  }))
  // ...
}
```
As you notice we need to use the martini instance as well.

To specify the template you want to display for your route, do the following:

``` go
func main() {
  // ...
  m.Get("/", func(r render.Render) {
    r.HTML(200, "index", nil)
  })
  // ...
}
```

Now you can render you view on the root path of your application.

If you notice we are are sending three parameters to the `m.HTML(...)`
method, and the order is the following: `m.HTML(statusCode, templateName, object)`.

In order to pass an object to show in the view you need to do the following:

``` go
import (
  // ...
)

type Person struct {
  Id      int
  Name    string
  Age     int
  Email   string
  Job     string
}

func main() {
  // ...
}
```
The struct will be the object we're going to send to the route view.

``` go
func main() {
  // ...
  m.Get("/", func(r render.Render) {
    person := Person{ 1, "John Doe", 33, "email@example.com", "Cashier" }
    r.HTML(200, "index", person)
  })
  // ...
}
```

And you can access it in the view like this:

``` mustache
<h1>Displaying {{ .Name }}</h1>

<div>
  <p>Id: {{ .Id}}</p>
  <p>Age: {{ .Age}}</p>
  <p>Email: {{ .Email}}</p>
  <p>Job: {{ .Job}}</p>
</div>
```

To more information about the Go templating engine please refer to [the documentation](http://golang.org/pkg/text/template/).

## Reading data from a postgresql database

Comming soon

## Inserting data to a postgresql database

Comming soon

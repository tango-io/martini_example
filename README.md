# Martini Example

This application is a simple CRUD application written in [Golang](http://golang.org/) using [Martini](http://martini.codegangsta.io/) and [PostgreSQL](http://www.postgresql.org/).


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

## Writing routes

## Rendering views

## Reading data from a postgresql database

## Inserting data to a postgresql database

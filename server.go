package main

import (
	"log"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func doMain(_ *log.Logger, r render.Render) {

	r.HTML(200, "index", nil)
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Extensions: []string{".tmpl"},
	}))

	// Router
	m.Get("/", doMain)

	m.Run()
}

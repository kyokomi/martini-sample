package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func doStarChart(r render.Render) {
	r.HTML(200, "index", nil)
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Extensions: []string{".tmpl"},
	}))

	// Router
	m.Get("/", doStarChart)

	m.Run()
}

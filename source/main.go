package main

import(
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func main() {

	m := martini.Classic()

	m.Use(render.Renderer())

	m.NotFound(func (r render.Render){
		r.Redirect("/")
	})

	m.Get("/", top)
	m.Get("/happy", happy)

	m.Run()
}

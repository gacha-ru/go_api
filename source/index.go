package main

import "github.com/martini-contrib/render"

func top(r render.Render) {
	r.JSON(200, map[string]interface{}{"hello": "world"})
}

func happy(r render.Render) {
	r.JSON(200, map[string]interface{}{"happy": "turn"})
}

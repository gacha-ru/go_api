package main

import (
	"github.com/martini-contrib/render"
	"github.com/garyburd/redigo/redis"
)

func top(r render.Render) {
	r.JSON(200, map[string]interface{}{"hello": "world"})
}

func happy(r render.Render) {
	r.JSON(200, map[string]interface{}{"happy": "turn"})
}

func religo(r render.Render) {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer  c.Close()

	count := 10

	// set
	c.Do("SET", "message", count)
	c.Do("INCR", "message")

	//get
	world, err := redis.String(c.Do("GET", "message"))
	if err != nil {
		r.HTML(500, "key not found","")
	}

	r.JSON(200, map[string]interface{}{"keys":world})
}

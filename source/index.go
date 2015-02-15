package main

import (
	"github.com/martini-contrib/render"
	"github.com/garyburd/redigo/redis"
	"github.com/go-martini/martini"
)

func top(r render.Render) {
	r.JSON(200, map[string]interface{}{"hello": "world"})
}

func happy(r render.Render) {
	r.JSON(200, map[string]interface{}{"happy": "turn"})
}

func religo(r render.Render, pool *redis.Pool, params martini.Params) {

	c := pool.Get()

	//get
	count, acerr := redis.String(c.Do("GET", "access_count"))
	if acerr != nil {
		count := 1
		c.Do("SET", "access_count", count)
	}

	// set
	c.Do("INCR", "access_count")
	c.Do("SET", "message", count)

	r.JSON(200, map[string]interface{}{"keys":count})
}

package main

import (
	"github.com/martini-contrib/render"
	"github.com/garyburd/redigo/redis"
	"github.com/go-martini/martini"
	"fmt"
	"net/http"
)

func setkey(r render.Render, pool *redis.Pool, params martini.Params, req *http.Request) {
	key := params["key"]
	value := req.URL.Query().Get("value")

	c := pool.Get()
	defer c.Close()

	status, err := c.Do("SET", key, value)

	if err != nil {
		message := fmt.Sprintf("Could not SET %s:%s", key, value)

		r.JSON(400, map[string]interface{}{
			"status":  "ERR",
			"message": message})
	} else {
		r.JSON(200, map[string]interface{}{
			"status": status})
	}
}

func getkey(r render.Render, pool *redis.Pool, params martini.Params) {
	key := params["key"]

	c := pool.Get()
	defer c.Close()

	value, err := redis.String(c.Do("GET", key))

	if err != nil {
		message := fmt.Sprintf("Could not GET %s", key)
		r.JSON(400, map[string]interface{}{
		"status":  "ERR",
		"message": message})
	} else {
		r.JSON(200, map[string]interface{}{
		"status": "OK",
		"value":  value})
	}
}

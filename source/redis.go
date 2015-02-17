package main

import (
	"github.com/martini-contrib/render"
	"github.com/garyburd/redigo/redis"
	"github.com/go-martini/martini"
	"fmt"
	"net/http"
)


// redisの値をインクリメントして表示
func religo(r render.Render, pool *redis.Pool, params martini.Params) {
	c := pool.Get()

	// set
	c.Do("INCR", "access_count")

	//get
	count, acerr := redis.String(c.Do("GET", "access_count"))
	if acerr != nil {
		count := 1
		c.Do("SET", "access_count", count)
	}

	r.JSON(200, map[string]interface{}{"keys":count})
}


func setkey(r render.Render, pool *redis.Pool, params martini.Params, req *http.Request) {
	key := params["key"]
	value := params["value"]
	//value := req.URL.Query().Get("value")

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

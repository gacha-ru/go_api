package main

import(
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/garyburd/redigo/redis"
	"flag")

var (
	redisAddress   = flag.String("redis-address", ":6379", "Address to the Redis server")
	maxConnections = flag.Int("max-connections", 10, "Max connections to Redis")
)

func main() {

	flag.Parse()

	redisPool := redis.NewPool(func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", *redisAddress)

		if err != nil {
			return nil, err
		}

		return c, err
	}, *maxConnections)

	defer redisPool.Close()

	m := martini.Classic()

	m.Use(render.Renderer())

	m.NotFound(func (r render.Render){
		r.Redirect("/")
	})

	m.Map(redisPool)

	m.Get("/", top)
	m.Get("/happy", happy)
	m.Get("/religo", religo)
	m.Get("/set/:key", setkey)
	m.Get("/get/:key", getkey)

	m.Run()
}

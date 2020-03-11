package main

import (
	"github.com/TonyXMH/gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>hello gee<h1>")
	})
	r.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you are at:%s\n", c.Query("name"), c.Path)
	})
	r.POST("/login", func(c *gee.Context) {
		c.Josn(http.StatusOK, gee.H{
			"username": c.PostFrom("username"),
			"password": c.PostFrom("password"),
		})
	})
	r.GET("/hello/:name", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})
	r.GET("/assets/*filepath", func(c *gee.Context) {
		c.Josn(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}

//curl http://localhost:9999/
//curl http://localhost:9999/hello
//curl http://localhost:9999/login -X POST -d 'username=tonyxin&password=123456'
//curl "http://localhost:9999/hello/tony"
//curl "http://localhost:9999/assets/css/picture.css"
package main

import (
	"github.com/TonyXMH/gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.Use(gee.Logger())
	r.GET("/index", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gee.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})
		v1.GET("/hello", func(c *gee.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})
	}

	v2 := r.Group("/v2")

	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			c.String(http.StatusOK, "hello %s, you are at:%s\n", c.Query("name"), c.Path)
		})
		v2.POST("/login", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}

	r.GET("/hello/:name", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})
	r.GET("/assets/*filepath", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}

//curl http://localhost:9999/
//curl http://localhost:9999/hello
//curl http://localhost:9999/login -X POST -d 'username=tonyxin&password=123456'
//curl "http://localhost:9999/hello/tony"
//curl "http://localhost:9999/assets/css/picture.css"

//curl "http://localhost:9999/index"
//curl "http://localhost:9999/v1/"
//curl "http://localhost:9999/v1/hello"
//curl "http://localhost:9999/v2/hello/tony"
//curl "http://localhost:9999/v2/login" -X POST -d 'username=tony&password=1234'

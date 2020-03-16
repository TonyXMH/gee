package main

import (
	"fmt"
	"github.com/TonyXMH/gee"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println("++++++++++", dir)
	r := gee.New()
	r.Use(gee.Logger())
	r.SetFuncMap(template.FuncMap{"formatAsDate": formatAsDate})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "static")
	stu1 := &student{
		Name: "tony",
		Age:  27,
	}
	stu2 := &student{
		Name: "tom",
		Age:  10,
	}
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *gee.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gee.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})
	r.GET("/date", func(c *gee.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", gee.H{
			"title": "gee",
			"now":   time.Now(),
		})
	})
	//
	//v1 := r.Group("/v1")
	//{
	//	v1.GET("/", func(c *gee.Context) {
	//		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	//	})
	//	v1.GET("/hello", func(c *gee.Context) {
	//		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	//	})
	//}
	//
	//v2 := r.Group("/v2")
	//
	//{
	//	v2.GET("/hello/:name", func(c *gee.Context) {
	//		c.String(http.StatusOK, "hello %s, you are at:%s\n", c.Query("name"), c.Path)
	//	})
	//	v2.POST("/login", func(c *gee.Context) {
	//		c.JSON(http.StatusOK, gee.H{
	//			"username": c.PostForm("username"),
	//			"password": c.PostForm("password"),
	//		})
	//	})
	//}
	//
	//r.GET("/hello/:name", func(c *gee.Context) {
	//	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	//})
	//r.GET("/assets/*filepath", func(c *gee.Context) {
	//	c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
	//})

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

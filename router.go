package gee

import (
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{make(map[string]HandlerFunc)}
}

func (r *router) addRoute(method, pattern string, handlerFunc HandlerFunc) {
	key := method + "-" + pattern
	r.handlers[key] = handlerFunc
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	//log.Printf("handlers:%+v",r.handlers)
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 not found:%s\n", c.Path)
	}
}

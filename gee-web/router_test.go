package gee_web

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_parsePattern(t *testing.T) {
	type args struct {
		pattern string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			"/p/:name",
			args{
				"/p/:name",
			},
			[]string{"p", ":name"},
		},
		{
			"/p/*",
			args{
				"/p/*",
			},
			[]string{"p", "*"},
		},
		{
			"/p/*name/*",
			args{
				"/p/*name/*",
			},
			[]string{"p", "*name"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parsePattern(tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parsePattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRoute("GET", "/hello/tony")
	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}
	if n.pattern != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}
	if ps["name"] != "tony" {
		t.Fatal("name should be equal to 'tony'")
	}
	fmt.Printf("matched path: %s, params['name']: %s\n", n.pattern, ps["name"])
}

func TestGetRoute2(t *testing.T) {
	r := newTestRouter()
	n1, ps1 := r.getRoute("GET", "/assets/file1.txt")
	ok1 := n1.pattern == "/assets/*filepath" && ps1["filepath"] == "file1.txt"
	if !ok1 {
		t.Fatal("pattern should be /assets/*filepath & filepath shoule be file1.txt")
	}
	n2, ps2 := r.getRoute("GET", "/assets/css/test.css")
	ok2 := n2.pattern == "/assets/*filepath" && ps2["filepath"] == "css/test.css"
	if !ok2 {
		t.Fatal("pattern shoule be /assets/*filepath & filepath shoule be css/test.css")
	}
}

func TestGetRoutes(t *testing.T) {
	r := newTestRouter()
	nodes := r.getRoutes("GET")
	for i, n := range nodes {
		fmt.Println(i+1, n)
	}
	if len(nodes) != 5 {
		t.Fatal("the number of routes shoule be 4")
	}
}

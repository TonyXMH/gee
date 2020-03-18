package main

import (
	"fmt"
	gee_cache "github.com/TonyXMH/gee/gee-cache"
	"log"
	"net/http"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func main() {
	gee_cache.NewGroup("scores", 2<<10, gee_cache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))
	addr := "localhost:9999"
	peers := gee_cache.NewHTTPPool(addr)
	log.Println("gee_cache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}

//curl http://localhost:9999/_geecache/scores/Jack
//curl http://localhost:9999/_geecache/scores/Tom

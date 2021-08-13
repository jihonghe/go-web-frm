package main

import (
	"core"
	"fmt"
	"log"
	"net/http"
)

// 活下去，寻找希望

func main() {
	router := core.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	})
	router.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	log.Fatal(router.Run("0.0.0.0:9999"))
}

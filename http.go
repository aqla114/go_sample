package main

import (
	"fmt"
	"net/http"
)

// IndexHandler requestを扱う。
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

// NewIndexHandler requestを扱う。
func NewIndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello new world")
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/new", NewIndexHandler)

	http.ListenAndServe(":3000", nil)
}

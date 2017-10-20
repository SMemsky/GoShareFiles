package main

import (
	"fmt"
	"net/http"
)

var fileHandler = http.FileServer(http.Dir("./"))

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	fileHandler.ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Starting...")
	panic(http.ListenAndServe(":65535", nil))
}

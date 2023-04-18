package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s: 200", r.URL.Path)
		fmt.Fprintf(w, "Hello, World!")
	})

	http.ListenAndServe(":8080", nil)
}

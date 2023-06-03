package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			fmt.Fprintf(w, "Hello, %q, This is POST with body: %q", "ChuanGz", r.Body)
		} else {

			fmt.Fprintf(w, "Hello, %q, This is %q", "ChuanGz", r.Method)
		}
	})
	log.Fatal(http.ListenAndServe(":3000", nil))
}

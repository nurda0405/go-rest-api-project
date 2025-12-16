package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})
	http.HandleFunc("/teachers", func(w http.ResponseWriter, r *http.Request) {

	})
	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			query := r.URL.Query()
			w.Write([]byte(query.Get("name")))
		}
	})
	http.HandleFunc("/execs", func(w http.ResponseWriter, r *http.Request) {

	})

	port := ":3000"
	fmt.Println("Server running on port", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error running the server:", err)
	}

}

package router

import (
	"net/http"
	"restapi/internal/api/handlers"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})
	mux.HandleFunc("/teachers/", handlers.TeachersHandler)

	mux.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {

	})
	mux.HandleFunc("/execs", func(w http.ResponseWriter, r *http.Request) {

	})
	return mux
}

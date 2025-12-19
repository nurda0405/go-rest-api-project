package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	mw "restapi/internal/api/middlewares"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})
	mux.HandleFunc("/teachers", func(w http.ResponseWriter, r *http.Request) {

	})
	mux.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			query := r.URL.Query()
			w.Write([]byte(query.Get("name")))
		}
	})
	mux.HandleFunc("/execs", func(w http.ResponseWriter, r *http.Request) {

	})
	cert := "cert.pem"
	key := "key.pem"
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	rl := mw.NewRateLimiter(2, 5*time.Second)

	port := ":3000"
	fmt.Println("Server running on port", port)
	server := &http.Server{
		Addr:      port,
		Handler:   rl.RateLimiterMiddleware(mw.Compression(mw.ResponseTimeMiddleware(mw.SecurityHeaders(mw.Cors(mux))))),
		TLSConfig: tlsConfig,
	}

	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error running the server:", err)
	}

}

package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	mw "restapi/internal/api/middlewares"
	"restapi/internal/api/router"
)

func main() {

	cert := "cert.pem"
	key := "key.pem"
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// rl := mw.NewRateLimiter(2, 5*time.Second)
	// hppOptions := mw.HPPOptions{
	// 	CheckBody:               true,
	// 	CheckQuery:              true,
	// 	CheckForOnlyContentType: "x-www-form-urlencoded",
	// 	Whitelist:               []string{"sortOrder", "sortBy", "name", "age", "class"},
	// }

	// cors rate time security compressioon hpp
	// secureMux := mw.Cors(rl.RateLimiterMiddleware(mw.ResponseTimeMiddleware(mw.SecurityHeaders(mw.Compression(mw.HPP(hppOptions)(mux))))))
	// secureMux := applyMiddlewares(mux, mw.HPP(hppOptions), mw.Compression, mw.SecurityHeaders, mw.ResponseTimeMiddleware, rl.RateLimiterMiddleware, mw.Cors)
	secureMux := mw.SecurityHeaders(router.Router())
	port := ":3000"
	fmt.Println("Server running on port", port)
	server := &http.Server{
		Addr:      port,
		Handler:   secureMux,
		TLSConfig: tlsConfig,
	}

	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error running the server:", err)
	}
}

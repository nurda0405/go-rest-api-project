package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func ResponseTimeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrappedWriter := &responseWriter{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(wrappedWriter, r)
		duration := time.Since(start)
		fmt.Println("Execution duration", duration.Milliseconds())
	})
}

// this struct is needed to know the status code of the response
// by default there is no status field in http.ResponseWriter
// in this struct we are connecting the status code with the http.ResponseWriter
// this ensures synchronization of actual status code and status field to use it in the code
type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

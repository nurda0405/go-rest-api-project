package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	mw "restapi/internal/api/middlewares"
	"strconv"
	"strings"
	"sync"
)

type Teacher struct {
	ID        int
	FirstName string
	LastName  string
	Class     string
	Subject   string
}

var (
	teachers = make(map[int]Teacher)
	mutex    = &sync.Mutex{}
	nextID   = 1
)

func init() {
	teachers[nextID] = Teacher{
		ID:        nextID,
		FirstName: "Nurlybai",
		LastName:  "Uzakbayev",
		Class:     "12D",
		Subject:   "Computer Science",
	}
	nextID++

	teachers[nextID] = Teacher{
		ID:        nextID,
		FirstName: "Anuar",
		LastName:  "N/A",
		Class:     "8D",
		Subject:   "Math",
	}
	nextID++

	teachers[nextID] = Teacher{
		ID:        nextID,
		FirstName: "Altynai",
		LastName:  "N/A",
		Class:     "12D",
		Subject:   "Kazakh Language",
	}
}
func getTeachersHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/teachers/")
	idStr := strings.TrimSuffix(path, "/")
	if idStr == "" {
		firstName := r.URL.Query().Get("first_name")
		lastName := r.URL.Query().Get("last_name")

		teacherList := make([]Teacher, 0, len(teachers))
		for _, teacher := range teachers {
			if (firstName == "" || teacher.FirstName == firstName) && (lastName == "" || teacher.LastName == lastName) {
				teacherList = append(teacherList, teacher)
			}
		}
		response := struct {
			Status string    `json:"status"`
			Count  int       `json:"count"`
			Data   []Teacher `json:"data"`
		}{
			Status: "success",
			Count:  len(teacherList),
			Data:   teacherList,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {

		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println(err)
			return
		}

		teacher, exists := teachers[id]
		if !exists {
			http.Error(w, "Teacher not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(teacher)
	}
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTeachersHandler(w, r)
	}
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})
	mux.HandleFunc("/teachers/", teachersHandler)

	mux.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {

	})
	mux.HandleFunc("/execs", func(w http.ResponseWriter, r *http.Request) {

	})
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
	secureMux := mw.SecurityHeaders(mux)
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

type Middleware func(http.Handler) http.Handler

func ApplyMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"http-server/config"
	"http-server/handlers"
	"http-server/static"
	"log"
	"net/http"
)

func main() {
	cfg := config.NewConfig()

	router := mux.NewRouter()

	router.HandleFunc("/", handlers.RootHandler).Methods("GET")
	router.HandleFunc("/hello", handlers.HelloHandler).Methods("GET")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.HandlerFunc(static.ServeStaticFiles)))

	//fs := http.FileServer(http.Dir("static"))
	//router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	filePath := filepath.Base(r.URL.Path)
	//	contentType := mime.TypeByExtension(filepath.Ext(filePath))
	//	if contentType != "" {
	//		w.Header().Set("Content-Type", contentType)
	//	}
	//	fs.ServeHTTP(w, r)
	//})))

	fmt.Printf("Server is running on %s\n", cfg.ServerAddress)
	err := http.ListenAndServe(cfg.ServerAddress, router)
	if err != nil {
		log.Fatal("Error:", err)
	}
}

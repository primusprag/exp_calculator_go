package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir(".")))
	mux.HandleFunc("GET /server", middleware.middlewareLog(handleApp))

	server := http.Server{
		Handler: mux,
		Addr:    ":8080", //later port

	}

	server.ListenAndServe()
}

func handleApp(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server Online"))
}

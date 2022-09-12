package app

import (
	"log"
	"net/http"
)

func App() error {
	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	log.Println("liteServer is starting >>>")
	log.Printf("Running at %v\n", server.Addr)

	return server.ListenAndServe()
}

package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	port := ":8080"

	router := http.NewServeMux()

	// Path Parameter
	router.HandleFunc("/item/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		w.Write([]byte("Recieved request for item id : " + id))
	})

	s := &http.Server{
		Addr:           port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Starting server on Port%s\n", port)
	log.Fatal(s.ListenAndServe())
}

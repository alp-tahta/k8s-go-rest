package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := ":8080"

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			log.Println(err)
			http.Error(w, "Failed to get hostname", http.StatusInternalServerError)
			return
		}
		_, err = w.Write([]byte("Hostname : " + hostname + time.Now().String()))
		if err != nil {
			log.Println(err)
			http.Error(w, "Could not write data", http.StatusInternalServerError)
			return
		}
	})

	// Path Parameter
	router.HandleFunc("GET /item/{id}", func(w http.ResponseWriter, r *http.Request) {
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

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	port := ":8080"

	userAPI := "http://localhost:8080"

	router := http.NewServeMux()

	router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.HandleFunc("GET /greet/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		requestURL := fmt.Sprintf("%s/%s/%s", userAPI, "user", id)
		res, err := http.Get(requestURL)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if res.StatusCode != http.StatusOK {
			log.Printf("Response code from userAPI is %d\n", res.StatusCode)
			http.Error(w, fmt.Sprintf("Response Code From userAPI is %d\n", res.StatusCode), http.StatusInternalServerError)
			return
		}

		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = w.Write(append([]byte("Hi "), resBody...))
		if err != nil {
			log.Println(err)
			http.Error(w, "Could not write data", http.StatusInternalServerError)
			return
		}
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

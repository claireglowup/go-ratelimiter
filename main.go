package main

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/time/rate"
)

type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := Message{
		Status: "Successfull",
		Body:   "Hi! you've reached the API, can i help you?",
	}
	log.Println(r.URL.Path)
	if err := json.NewEncoder(w).Encode(&message); err != nil {
		return
	}
}

func rateLimiter(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	limiter := rate.NewLimiter(2, 4)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {

			message := Message{
				Status: "Request Failed",
				Body:   "The API is at capacity, try again later",
			}
			log.Println(message.Body)
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(message)
			return
		} else {
			next(w, r)
		}
	})
}

func main() {
	http.Handle("/ping", rateLimiter(handler))
	log.Println("server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err.Error())
	}
}

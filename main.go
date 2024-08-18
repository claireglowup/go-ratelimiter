package main

import (
	"encoding/json"
	"log"
	"net/http"
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
	if err := json.NewEncoder(w).Encode(&message); err != nil {
		return
	}
}

func main() {
	http.Handle("/ping", rateLimiter(handler))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err.Error())
	}
}

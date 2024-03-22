package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/welcome", welcome)
	fmt.Println("Starting HTTP server....")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Println(err)
	}
}

func welcome(w http.ResponseWriter, r *http.Request) {
	res := Response{
		Status:  true,
		Message: "Request processed successfully!",
		Data:    "Welcome to tech trainings ",
	}

	// Convert struct to bytes
	reInBytes, err := json.Marshal(res)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}

	// Print results
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(reInBytes)
}

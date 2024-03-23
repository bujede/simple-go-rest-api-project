package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Tasks - a slice to store our tasks in-memory
var tasks = []Task{}

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Task - Define a simple struct for our example
type Task struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/welcome", welcome)
	// Route handles & endpoints
	r.HandleFunc("/tasks", getTasks).Methods("GET")
	r.HandleFunc("/task", CreateTask).Methods("POST")
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

// createTask adds a new task to our in-memory store.
func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tasks = append(tasks, task) // Add the task to the in-memory store
	json.NewEncoder(w).Encode(task)
}

// getTasks responds with the list of all tasks as JSON.
func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

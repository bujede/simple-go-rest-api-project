package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestWelcomeHandler tests the welcome handler
func TestWelcomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/welcome", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(welcome)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"status":true,"message":"Request processed successfully!","data":"Welcome to tech trainings "}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// TestCreateTask tests the CreateTask function
func TestCreateTask(t *testing.T) {
	var jsonStr = []byte(`{"id":"1", "name":"Test Task", "description":"This is a test task."}`)
	req, err := http.NewRequest("POST", "/task", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTask)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var task Task
	if err := json.Unmarshal(rr.Body.Bytes(), &task); err != nil {
		t.Fatal(err)
	}

	if task.ID != "1" || task.Name != "Test Task" {
		t.Errorf("handler returned unexpected body: got %v", rr.Body.String())
	}
}

// TestGetTasks tests the getTasks function
func TestGetTasks(t *testing.T) {
	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getTasks)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Expecting at least one task in the response after the TestCreateTask runs
	// This assumes tests are run sequentially which is the default behavior
	if rr.Body.String() == "[]" {
		t.Error("handler returned an empty array, expected at least one task")
	}
}

package routes

import (
	"encoding/json"
	"main/chatbot/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProducRoutes(t *testing.T) {
	// Create a request to the "/api/" route
	req := httptest.NewRequest("GET", "/api/", nil)
	rr := httptest.NewRecorder()

	router := ProducRoutes()
	router.ServeHTTP(rr, req)

	// Check the status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	// Parse the response body
	var message models.Message
	err := json.NewDecoder(rr.Body).Decode(&message)
	if err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	// Check the response message
	expectedMessage := "Store API"
	if message.Message != expectedMessage {
		t.Errorf("Expected message: %s, but got: %s", expectedMessage, message.Message)
	}
}

package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// Helper function to set up the router for testing
func setupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("v1")
	{
		v1.GET("/risks", getRisks)
		v1.POST("/risks", createRisk)
		v1.GET("/risks/:id", getRiskByID)
	}
	return r
}

// Test GET /v1/risks (when there are no risks yet)
func TestGetRisksEmpty(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/v1/risks", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, w.Code)
	}

	expected := "[]"
	if w.Body.String() != expected {
		t.Errorf("Expected body %v but got %v", expected, w.Body.String())
	}
}

// Test POST /v1/risks (successful creation)
func TestCreateRisk(t *testing.T) {
	router := setupRouter()

	risk := Risk{
		State:       "open",
		Title:       "Test Risk",
		Description: "This is a test risk",
	}
	body, _ := json.Marshal(risk)

	req, _ := http.NewRequest("POST", "/v1/risks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status %d but got %d", http.StatusCreated, w.Code)
	}

	var createdRisk Risk
	json.NewDecoder(w.Body).Decode(&createdRisk)

	if createdRisk.ID == "" {
		t.Errorf("Expected a valid UUID, but got empty ID")
	}

	if createdRisk.State != risk.State || createdRisk.Title != risk.Title || createdRisk.Description != risk.Description {
		t.Errorf("Expected risk %v but got %v", risk, createdRisk)
	}
}

// Test GET /v1/risks/{id} (retrieving a specific risk)
func TestGetRiskByID(t *testing.T) {
	router := setupRouter()

	// First, create a risk
	risk := Risk{
		State:       "open",
		Title:       "Test Risk",
		Description: "This is a test risk",
	}
	body, _ := json.Marshal(risk)

	req, _ := http.NewRequest("POST", "/v1/risks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status %d but got %d", http.StatusCreated, w.Code)
	}

	var createdRisk Risk
	json.NewDecoder(w.Body).Decode(&createdRisk)

	// Now, fetch the risk by ID
	req, _ = http.NewRequest("GET", "/v1/risks/"+createdRisk.ID, nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, w.Code)
	}

	var fetchedRisk Risk
	json.NewDecoder(w.Body).Decode(&fetchedRisk)

	if fetchedRisk.ID != createdRisk.ID {
		t.Errorf("Expected risk ID %v but got %v", createdRisk.ID, fetchedRisk.ID)
	}
}

// Test GET /v1/risks/{id} for a non-existent risk
func TestGetRiskByIDNotFound(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/v1/risks/nonexistent-id", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status %d but got %d", http.StatusNotFound, w.Code)
	}

	expected := `{"error":"Risk not found"}`
	if w.Body.String() != expected {
		t.Errorf("Expected body %v but got %v", expected, w.Body.String())
	}
}

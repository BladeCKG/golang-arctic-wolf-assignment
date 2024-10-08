package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Define the structure for a Risk
type Risk struct {
	ID          string `json:"id"`
	State       string `json:"state"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Define a custom type for the RiskState
type RiskState string

// Define constants for the possible values
const (
	Open          RiskState = "open"
	Closed        RiskState = "closed"
	Accepted      RiskState = "accepted"
	Investigating RiskState = "investigating"
)

// Function to return all possible RiskState values
func AllRiskStates() []RiskState {
	return []RiskState{
		Open,
		Closed,
		Accepted,
		Investigating,
	}
}

// In-memory store for risks
var (
	riskStore = make(map[string]Risk)
	mutex     = &sync.Mutex{} // To protect concurrent access to the riskStore
)

// Get all risks
func getRisks(c *gin.Context) {
	mutex.Lock()
	defer mutex.Unlock()

	risks := []Risk{}
	for _, risk := range riskStore {
		risks = append(risks, risk)
	}

	c.JSON(http.StatusOK, risks)
}

// Create a new risk
func createRisk(c *gin.Context) {
	var risk Risk
	if err := c.ShouldBindJSON(&risk); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if risk.State == "" || risk.Title == "" || risk.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required risk fields"})
		return
	}

	isCorrectState := false
	for _, state := range AllRiskStates() {
		if state == RiskState(risk.State) {
			isCorrectState = true
			break
		}
	}

	if !isCorrectState {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Risk State"})
		return
	}

	// Generate a UUID for the new risk
	risk.ID = uuid.New().String()

	// Save the risk in the in-memory store
	mutex.Lock()
	riskStore[risk.ID] = risk
	mutex.Unlock()

	c.JSON(http.StatusCreated, risk)
}

// Get a specific risk by ID
func getRiskByID(c *gin.Context) {
	id := c.Param("id")

	mutex.Lock()
	risk, exists := riskStore[id]
	mutex.Unlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Risk not found"})
		return
	}

	c.JSON(http.StatusOK, risk)
}

func main() {
	// Initialize Gin router
	router := gin.Default()

	// Version 1 of the API
	v1 := router.Group("/v1")
	{
		// Define the endpoints
		v1.GET("/risks", getRisks)
		v1.POST("/risks", createRisk)
		v1.GET("/risks/:id", getRiskByID)
	}

	// Start the server
	router.Run(":8080")
}

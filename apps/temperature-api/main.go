package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// TemperatureResponse represents the response structure
type TemperatureResponse struct {
	Value       float64   `json:"value"`
	Unit        string    `json:"unit"`
	Timestamp   time.Time `json:"timestamp"`
	Location    string    `json:"location"`
	Status      string    `json:"status"`
	SensorID    string    `json:"sensor_id"`
	SensorType  string    `json:"sensor_type"`
	Description string    `json:"description"`
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main() {
	rand.Seed(time.Now().UnixNano())

	router := gin.Default()

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Temperature endpoint for location query
	router.GET("/temperature", func(c *gin.Context) {
		location := c.Query("location")
		sensorID := ""

		// If no location is provided, use a default based on sensor ID (but since sensorID is not in query, set to empty)
		if location == "" {
			location = "Unknown"
		}

		// If no sensor ID is provided, generate one based on location
		switch location {
		case "Living Room":
			sensorID = "1"
		case "Bedroom":
			sensorID = "2"
		case "Kitchen":
			sensorID = "3"
		default:
			sensorID = "0"
		}

		// Generate random temperature between 15 and 30
		value := 15 + rand.Float64()*15

		response := TemperatureResponse{
			Value:       value,
			Unit:        "Celsius",
			Timestamp:   time.Now(),
			Location:    location,
			Status:      "online",
			SensorID:    sensorID,
			SensorType:  "temperature",
			Description: "Current temperature reading",
		}

		c.JSON(http.StatusOK, response)
	})

	// Temperature endpoint for sensor ID
	router.GET("/temperature/:sensorID", func(c *gin.Context) {
		sensorID := c.Param("sensorID")
		location := ""

		// If no location is provided, use a default based on sensor ID
		switch sensorID {
		case "1":
			location = "Living Room"
		case "2":
			location = "Bedroom"
		case "3":
			location = "Kitchen"
		default:
			location = "Unknown"
		}

		// Generate random temperature between 15 and 30
		value := 15 + rand.Float64()*15

		response := TemperatureResponse{
			Value:       value,
			Unit:        "Celsius",
			Timestamp:   time.Now(),
			Location:    location,
			Status:      "online",
			SensorID:    sensorID,
			SensorType:  "temperature",
			Description: "Current temperature reading",
		}

		c.JSON(http.StatusOK, response)
	})

	port := getEnv("PORT", ":8081")
	log.Printf("Starting temperature-api on port %s", port)
	if err := router.Run(port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

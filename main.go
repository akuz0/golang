package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// flight represents data about a record flight.
type flight struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Destination string  `json:"destination"`
	Price       float64 `json:"price"`
}

// flights slice to seed record flight data.
var flights = []flight{
	{ID: "1", Name: "Viktor Korobkov", Destination: "CLT", Price: 56.99},
	{ID: "2", Name: "Petr Kuznecov", Destination: "YUL", Price: 17.99},
	{ID: "3", Name: "Vika Odyntsova", Destination: "JFK", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/flights", getFlights)
	router.GET("/flights/:id", getFlightsByID)
	router.POST("/flights", postFlight)

	router.Run("localhost:8091")
}

// getFlights responds with the list of all flights as JSON.
func getFlights(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, flights)
}

// postFlight adds an flight from JSON received in the request body.
func postFlight(c *gin.Context) {
	var newFlight flight

	// Call BindJSON to bind the received JSON to
	// newFlight.
	if err := c.BindJSON(&newFlight); err != nil {
		return
	}

	// Add the new flight to the slice.
	flights = append(flights, newFlight)
	c.IndentedJSON(http.StatusCreated, newFlight)
}

// getFlightsByID locates the flight whose ID value matches the id
// parameter sent by the client, then returns that flight as a response.
func getFlightsByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of flights, looking for
	// an flight whose ID value matches the parameter.
	for _, a := range flights {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "flight not found"})
}

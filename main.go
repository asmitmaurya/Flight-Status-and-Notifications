package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Flight struct represents the structure of a flight
type Flight struct {
	FlightNumber string `json:"flight_number"`
	Status       string `json:"status"`
	Gate         string `json:"gate"`
}

// Mock data for flights (replace with actual data fetching from airport system)
var flights = []Flight{
	{FlightNumber: "ABC123", Status: "On time", Gate: "A1"},
	{FlightNumber: "XYZ789", Status: "On time", Gate: "B2"},
}

// GetFlightsHandler returns all flights
func GetFlightsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flights)
}

// UpdateStatusHandler updates the status of a specific flight
func UpdateStatusHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	flightNumber := params["flightNumber"]

	var updatedStatus struct {
		Status string `json:"status"`
	}
	err := json.NewDecoder(r.Body).Decode(&updatedStatus)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i := range flights {
		if flights[i].FlightNumber == flightNumber {
			flights[i].Status = updatedStatus.Status
			// Here you can integrate notification logic if needed
			json.NewEncoder(w).Encode(flights[i])
			return
		}
	}

	http.Error(w, "Flight not found", http.StatusNotFound)
}

func main() {
	// Initialize Gorilla Mux router
	router := mux.NewRouter()

	// API endpoints
	router.HandleFunc("/api/flights", GetFlightsHandler).Methods("GET")
	router.HandleFunc("/api/update_status/{flightNumber}", UpdateStatusHandler).Methods("PUT")

	// Start server
	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

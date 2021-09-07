package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/noisersup/dashboard-backend-motivational/models"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/quote", getQuote).Methods("GET")

	log.Printf("Starting http server on port :8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getQuote(w http.ResponseWriter, r *http.Request) {
	log.Print("GET")

	response := models.GetResponse{}
	response.Quote = "Ledu." //TODO: Get from database

	sendResponse(w, response, http.StatusOK)
}

func sendResponse(w http.ResponseWriter, response interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("JSON encoding error: %s", err) //TODO: Log file
	}
}

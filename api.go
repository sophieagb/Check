package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ValidationRequest struct {
	CardNumber string `json:"card_number"`
}

type ValidationResponse struct {
	IsValid bool `json:"is_valid"`
}

func validateCardHandler(w http.ResponseWriter, r *http.Request) {
	var request ValidationRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// call the LuhnCheck function 
	isValid := LuhnCheck(request.CardNumber)

	response := ValidationResponse{IsValid: isValid}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/validate-card", validateCardHandler)
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}

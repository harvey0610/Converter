package api

import (
	"encoding/json"
	"net/http"
)

type ConversionRequest struct {
	Amount float64 `json:"amount"`
	From   string  `json:"from"`
	To     string  `json:"to"`
}

type ConversionResponse struct {
	Converted float64 `json:"converted"`
}

func ConvertHandler(w http.ResponseWriter, r *http.Request) {
	var req ConversionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// ⚠️ Mock exchange rates — for now we simulate real rates
	rates := map[string]float64{
		"USD": 1.0,
		"PHP": 56.82,
		"EUR": 0.93,
		"JPY": 155.24,
	}

	fromRate, ok1 := rates[req.From]
	toRate, ok2 := rates[req.To]

	if !ok1 || !ok2 {
		http.Error(w, "Invalid currency code", http.StatusBadRequest)
		return
	}

	converted := req.Amount / fromRate * toRate

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ConversionResponse{Converted: converted})
}

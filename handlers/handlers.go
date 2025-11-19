package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/uyendava/go-calculator/models"
)

// CalculateHandler handles /calculate POST requests
func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(models.CalcResponse{Message: "Only POST method is allowed"})
		return
	}

	var req models.CalcRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.CalcResponse{Message: "Invalid JSON input"})
		return
	}

	var result float64
	switch req.Operator {
	case "+":
		result = req.Num1 + req.Num2
	case "-":
		result = req.Num1 - req.Num2
	case "*":
		result = req.Num1 * req.Num2
	case "/":
		if req.Num2 == 0 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.CalcResponse{Message: "Division by zero!"})
			return
		}
		result = req.Num1 / req.Num2
	default:
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.CalcResponse{Message: "Invalid operator! Use +, -, *, or /"})
		return
	}

	// Send JSON response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.CalcResponse{Result: result})
}

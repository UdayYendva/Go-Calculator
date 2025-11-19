package models

// CalcRequest represents the JSON input for calculation
type CalcRequest struct {
	Num1     float64 `json:"num1"`
	Num2     float64 `json:"num2"`
	Operator string  `json:"operator"`
}

// CalcResponse represents the JSON output
type CalcResponse struct {
	Result  float64 `json:"result,omitempty"`
	Message string  `json:"message,omitempty"`
}

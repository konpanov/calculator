package main

import (
	"calculator/services"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", healthHandler)
	mux.HandleFunc("POST /calculate", calcuateHandler)

	handler := applyMiddleware(mux, []Middleware{corsMiddleware})
	log.Println("Server running on :9090")
	if err := http.ListenAndServe(":9090", handler); err != nil {
		log.Fatal(err)
	}
}

type Middleware func(http.Handler) http.Handler

func applyMiddleware(handler http.Handler, middlwares []Middleware) http.Handler {
	for _, middlware := range middlwares {
		handler = middlware(handler)
	}
	return handler
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

type CalculateInput struct {
	Expression string `json:"expression"`
}

type CalculateStatusCode int

const (
	CalculateStatusCodeSuccess CalculateStatusCode = iota
	CalculateStatusCodeInvalidInput
	CalculateStatusCodeDivisionByZero
	CalculateStatusCodeNegativeSqrt
)

type CalculateOutput struct {
	StatusCode CalculateStatusCode `json:"status_code"`
	Expression string              `json:"expression"`
	Result     string              `json:"result"`
	Error      string              `json:"error"`
}

func calcuateHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	var input CalculateInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Println(input)

	var output CalculateOutput
	output.Expression = input.Expression
	expressionResult, err := services.Eval(input.Expression)
	if err == nil {
		output.Result = fmt.Sprintf("%v", expressionResult)
		output.StatusCode = CalculateStatusCodeSuccess
	} else {
		output.Error = err.Error()
		switch err {
		case services.ErrorDevisionByZero:
			output.StatusCode = CalculateStatusCodeDivisionByZero
		case services.ErrorNegativeSqrt:
			output.StatusCode = CalculateStatusCodeNegativeSqrt
		default:
			w.WriteHeader(http.StatusBadRequest)
			output.StatusCode = CalculateStatusCodeInvalidInput
		}
	}

	if err := json.NewEncoder(w).Encode(output); err != nil {
		w.Write([]byte(err.Error()))
	}
}

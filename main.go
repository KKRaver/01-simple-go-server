package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/hello", helloWorld)
	r.HandleFunc("/hello2", headersMiddleware(helloWorldWithData))
	err := http.ListenAndServe(":8090", r)
	if err != nil {
		fmt.Println("Error occurred", err)
	}
}

type helloWorldResponse struct {
	Message string `json:"message"`
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	resp := helloWorldResponse{
		Message: "Hello World!",
	}
	jsonResp, err := json.Marshal(&resp)
	if err != nil {
		fmt.Println("Failed to parse response to JSON")
	}
	w.Write(jsonResp)
}

type helloWorldRequest struct {
	Name string `json:"name"`
}

func helloWorldWithData(w http.ResponseWriter, r *http.Request) {
	var req helloWorldRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println("Failed to decode incomign JSON", err)
		w.WriteHeader(400)
		return
	}
	w.WriteHeader(200)
	resp := helloWorldResponse{
		Message: "Hello " + req.Name,
	}
	jsonResp, err := json.Marshal(&resp)
	if err != nil {
		fmt.Println("Failed to parse response to JSON")
	}
	fmt.Println("Inside request")
	w.Write(jsonResp)
}

func headersMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Before request")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
		fmt.Println("After request")
	})
}

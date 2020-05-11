package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// M is interface of string
type M map[string]interface{}

// ActionData is method for endpoint
func ActionData(w http.ResponseWriter, r *http.Request) {
	log.Println("incoming request with method", r.Method)

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	payload := make(M)
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, ok := payload["Name"]; !ok {
		http.Error(w, "Payload `Name` is required", http.StatusBadRequest)
		return
	}

	if _, ok := payload["Age"]; !ok {
		http.Error(w, "Payload `Age` is required", http.StatusBadRequest)
		return
	}

	data := M{
		"Message": fmt.Sprintf("Hello %s, I am %s years old, I live in %s", payload["Name"], payload["Age"], payload["Address"]),
		"Status":  true,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	mux := new(http.ServeMux)
	mux.HandleFunc("/data", ActionData)

	server := new(http.Server)
	server.Handler = mux
	server.Addr = ":7000"

	log.Println("Starting server at", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln("Failed Start web server", err)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"log"
	"net/http"

	"github.com/ayush-yadavv/hospital_agents/types"
)

type Config struct {
	Agents []types.Personality `json:"agents"`
}

func main() {
	log.Println("Starting application...")

	// Load personalities
	log.Println("Loading personalities from JSON...")
	data, err := os.ReadFile("personalities.json")
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded %d personalities", len(config.Agents))

	// Create router and shared memory
	log.Println("Creating router...")
	router := NewRouter()

	// Create and register agents
	log.Println("Registering agents...")
	for _, personality := range config.Agents {
		pers := personality // Create new variable to avoid pointer issues
		log.Printf("Registering agent: %s (Role: %s)", pers.Name, pers.Role)
		agent := NewCustomerServiceAgent(&pers, router.memory)
		router.RegisterAgent(agent)
	}

	// Set up HTTP server
	http.HandleFunc("/message", router.HandleMessage)

	log.Println("Starting server on :8080...")
	fmt.Println("Server is running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ayush-yadavv/hospital_agents/types"
)

type Router struct {
	agents []types.Agent
	memory *types.Memory
}

// Add this struct for the response format
type AgentResponse struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

func NewRouter() *Router {
	return &Router{
		agents: make([]types.Agent, 0),
		memory: types.NewMemory(),
	}
}

func (r *Router) RegisterAgent(agent types.Agent) {
	r.agents = append(r.agents, agent)
}

func (r *Router) HandleMessage(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse JSON request
	var payload struct {
		Message string          `json:"message"`
		History []types.Message `json:"history"`
	}

	if err := json.NewDecoder(req.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if payload.Message == "" {
		http.Error(w, "Message is required", http.StatusBadRequest)
		return
	}

	// Initialize history if it's nil
	if payload.History == nil {
		payload.History = make([]types.Message, 0)
	}

	// Add user's message to history
	payload.History = append(payload.History, types.Message{
		From:    "User",
		Content: payload.Message,
	})

	// Find responding agents
	var respondingAgents []types.Agent
	for _, agent := range r.agents {
		for _, invocation := range agent.GetPersonality().Invocations {
			if strings.Contains(strings.ToLower(payload.Message), strings.ToLower(invocation)) {
				respondingAgents = append(respondingAgents, agent)
				break
			}
		}
	}

	responses := make([]AgentResponse, 0)
	currentHistory := payload.History

	// Process messages through each responding agent sequentially
	if len(respondingAgents) > 0 {
		for _, agent := range respondingAgents {
			response := agent.ProcessMessage(payload.Message, currentHistory)
			agentResp := AgentResponse{
				Name:    agent.GetPersonality().Name,
				Message: response,
			}
			responses = append(responses, agentResp)

			// Add this agent's response to history for next agent
			currentHistory = append(currentHistory, types.Message{
				From:    agent.GetPersonality().Name,
				Content: response,
			})
		}
	} else {
		// Default to default staff if no specific agents matched
		for _, agent := range r.agents {
			if agent.GetPersonality().Name == "default staff" {
				response := agent.ProcessMessage(payload.Message, currentHistory)
				responses = append(responses, AgentResponse{
					Name:    "default staff",
					Message: response,
				})
				break
			}
		}
	}

	// Return responses
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responses)
}

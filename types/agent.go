package types

type Agent interface {
	ProcessMessage(message string, history []Message) string
	GetPersonality() *Personality
}

type Personality struct {
	Name        string   `json:"name"`
	Role        string   `json:"role"`
	Personality string   `json:"personality"`
	Invocations []string `json:"invocations"`
}

type BaseAgent struct {
	Personality *Personality
	MemoryChan  chan string
}

type Message struct {
	From    string `json:"from"`
	Content string `json:"content"`
}

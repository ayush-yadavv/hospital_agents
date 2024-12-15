package types

type Memory struct {
	Channel chan string
}

func NewMemory() *Memory {
	return &Memory{
		Channel: make(chan string, 100), // buffered channel for memory sharing
	}
}

package main

// Message represents a single message in the conversation
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatHistory maintains the conversation history
type ChatHistory struct {
	messages []Message
}

// NewChatHistory creates a new chat history
func NewChatHistory() *ChatHistory {
	return &ChatHistory{
		messages: make([]Message, 0),
	}
}

// AddMessage adds a message to the history
func (ch *ChatHistory) AddMessage(role, content string) {
	ch.messages = append(ch.messages, Message{
		Role:    role,
		Content: content,
	})
}

// GetMessages returns all messages in the history
func (ch *ChatHistory) GetMessages() []Message {
	return ch.messages
}

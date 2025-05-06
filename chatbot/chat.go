package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

// Message represents a single message in the conversation
type Message struct {
	Role      string    `json:"role"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
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
		Role:      role,
		Content:   content,
		Timestamp: time.Now(),
	})
}

// GetMessages returns all messages in the history
func (ch *ChatHistory) GetMessages() []Message {
	return ch.messages
}

// SaveToFile saves the chat history to a JSON file
func (ch *ChatHistory) SaveToFile(filename string) error {
	data, err := json.MarshalIndent(ch.messages, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

// LoadFromFile loads chat history from a JSON file
func (ch *ChatHistory) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var messages []Message
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&messages); err != nil {
		return err
	}
	ch.messages = messages
	return nil
}

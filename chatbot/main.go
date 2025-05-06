package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	// Load configuration
	config, err := LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading configuration: %v\n", err)
		os.Exit(1)
	}

	// Initialize chat history and OpenAI client
	chatHistory := NewChatHistory()
	client := NewOpenAIClient(config)

	// Add system message to set the context
	chatHistory.AddMessage("system", "You are Chatty Charles, a helpful and friendly AI assistant. You provide clear, concise, and accurate responses while maintaining a conversational tone.")

	// Set up signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Create a scanner for reading user input
	scanner := bufio.NewScanner(os.Stdin)

	// Print welcome art
	printWelcomeArt()
	fmt.Printf("Using model: %s\n", config.Model)

	// Main conversation loop
	for {
		// Print prompt
		fmt.Print("\nYou: ")

		// Read user input
		if !scanner.Scan() {
			break
		}

		// Get user input and trim whitespace
		input := strings.TrimSpace(scanner.Text())

		// Check for exit command
		if strings.ToLower(input) == "exit" {
			printGoodbyeArt()
			return
		}

		// Add user message to history
		chatHistory.AddMessage("user", input)

		// Get AI response
		response, err := client.SendMessage(chatHistory.GetMessages())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting AI response: %v\n", err)
			continue
		}

		// Add AI response to history
		chatHistory.AddMessage("assistant", response)

		// Print AI response
		fmt.Printf("AI: %s\n", response)
	}

	// Handle any scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}

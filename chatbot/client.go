package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// OpenAIRequest represents the request body for the OpenAI API
type OpenAIRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
}

// OpenAIResponse represents the response from the OpenAI API
type OpenAIResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
	Error *struct {
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

// OpenAIClient handles communication with the OpenAI API
type OpenAIClient struct {
	apiKey      string
	model       string
	temperature float64
	httpClient  *http.Client
}

// NewOpenAIClient creates a new OpenAI client
func NewOpenAIClient(config *Config) *OpenAIClient {
	return &OpenAIClient{
		apiKey:      config.OpenAIAPIKey,
		model:       config.Model,
		temperature: config.Temperature,
		httpClient:  &http.Client{},
	}
}

// SendMessage sends a message to the OpenAI API and returns the response
func (c *OpenAIClient) SendMessage(messages []Message) (string, error) {
	// Prepare the request body
	reqBody := OpenAIRequest{
		Model:       c.model,
		Messages:    messages,
		Temperature: c.temperature,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %v", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	// Send the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}

	// Parse the response
	var openAIResp OpenAIResponse
	if err := json.Unmarshal(body, &openAIResp); err != nil {
		return "", fmt.Errorf("error parsing response: %v", err)
	}

	// Check for API errors
	if openAIResp.Error != nil {
		return "", fmt.Errorf("OpenAI API error: %s", openAIResp.Error.Message)
	}

	// Check if we got any choices
	if len(openAIResp.Choices) == 0 {
		return "", fmt.Errorf("no response from OpenAI API")
	}

	// Return the assistant's message
	return openAIResp.Choices[0].Message.Content, nil
}

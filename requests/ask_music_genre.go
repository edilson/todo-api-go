package requests

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"todo-api/helpers"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Messages            []Message   `json:"messages"`
	Model               string      `json:"model"`
	Temperature         float64     `json:"temperature"`
	MaxCompletionTokens int         `json:"max_completion_tokens"`
	TopP                float64     `json:"top_p"`
	Stream              bool        `json:"stream"`
	ReasoningEffort     string      `json:"reasoning_effort"`
	Stop                interface{} `json:"stop"`
}

var AskMusicGenre = func(todoTitle string) string {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY environment variable not set")
	}

	url := "https://api.groq.com/openai/v1/chat/completions"

	payload := ChatRequest{
		Messages: []Message{
			{Role: "user", Content: "Based on the task:" + todoTitle + "which type of music genre do you think i should listen during the time i'm doing this? Please narrow it down to one genre and return only the name of the genre on the response."},
		},
		Model:               "openai/gpt-oss-20b",
		Temperature:         1,
		MaxCompletionTokens: 8192,
		TopP:                1,
		Stream:              false,
		ReasoningEffort:     "medium",
		Stop:                nil,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Fatal("Error marshalling JSON: %v", err)
	}

	resp, err := helpers.PerformHttpRequest(url, apiKey, "POST", jsonData)

	fmt.Println("Genre returned", string(resp))

	var apiResponse struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	err = json.Unmarshal(resp, &apiResponse)
	if err != nil {
		log.Fatalf("Error unmarshalling response JSON: %v", err)
	}

	fmt.Println("API Response:", apiResponse.Choices[0].Message.Content)

	return apiResponse.Choices[0].Message.Content
}

package openai

import (
	"context"
	"fmt"

	"main/chatbot/db"

	"github.com/sashabaranov/go-openai"
)

func SendToGPT(apiKey, prompt string) string {
	client := openai.NewClient(apiKey)
	var messages []openai.ChatCompletionMessage
	for _, msg := range db.Conversation {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    msg.Role,
			Content: msg.Content,
		})
	}

	// Uncomment for debugging
	// for _, msg := range messages {
	// 	log.Printf("Role: %s", msg.Role)
	// 	log.Printf("Content: %s", msg.Content)
	// }
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       openai.GPT3Dot5Turbo,
			Messages:    messages,
			MaxTokens:   1024,
			Temperature: 1,
			TopP:        1,
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return fmt.Sprintf("ChatCompletion error: %v\n", err)
	}

	return resp.Choices[0].Message.Content
}

package chatbot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleBotUpdates(t *testing.T) {
	review := "This is a review without calling chatGPT"
	expectedResponse := respondDefault

	response := HandleBotUpdates(review)

	assert.Equal(t, response, expectedResponse)
}
package botClient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUpdate(t *testing.T) {
	var expected = UpdateConfig{
		Offset:1, 
		Limit:0, 
		Timeout:0, 
		AllowedUpdates:[]string(nil),
	}
	result := NewUpdate(1)
	assert.Equal(t, expected, result)
}

func TestNewMessage(t *testing.T) {
	var expected = MessageConfig{
		BaseChat: BaseChat{
			ChatID:           1,
			ReplyToMessageID: 0,
		},
		Text:                  "text",
		DisableWebPagePreview: false,
	}
	result := NewMessage(1, "text")
	assert.Equal(t, expected, result)
}


package botClient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	TestToken               = "12345"
)

func getBot(t *testing.T) (*BotAPI, error) {
	bot, err := NewBotAPI(TestToken)
	bot.Debug = true

	if err != nil {
		t.Error(err)
	}

	return bot, err
}

func TestNewBotAPI_notoken(t *testing.T) {
	res, err := NewBotAPI("")

	if err == nil {
		t.Error(err)
	}
	var nilBot = (*BotAPI)(nil)
	assert.Equal(t, nilBot, res)
}
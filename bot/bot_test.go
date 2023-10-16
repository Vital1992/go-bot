package bot

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	TestToken               = "6620296838:AAFGeDHI31nXwpxW0tObD0HwlUO9c1NyBro"
	ChatID                  = 667349672
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

func TestGetUpdates(t *testing.T) {
	bot, _ := getBot(t)

	u := NewUpdate(0)

	_, err := bot.GetUpdates(u)

	if err != nil {
		t.Error(err)
	}
}

func TestSendWithMessage(t *testing.T) {
	bot, _ := getBot(t)

	msg := NewMessage(ChatID, "A test message from the test library in telegram-bot-api")
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
	}
}

func ExampleNewBotAPI() {
	bot, err := NewBotAPI("MyAwesomeBotToken")
	if err != nil {
		panic(err)
	}

	u := NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Optional: wait for updates and clear them if you don't want to handle
	// a large backlog of old messages
	time.Sleep(time.Millisecond * 500)
	updates.Clear()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := NewMessage(update.Message.Chat.ID, update.Message.Text)

		bot.Send(msg)
	}
}
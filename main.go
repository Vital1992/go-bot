package main

import (
	"log"
	chatbot "main/bot"
	"os"
	// tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	//TODO: store key
	bot, err := chatbot.NewBotAPI(os.Getenv("API_KEY"))
	if err != nil {
		log.Panic(err)
	}

	// var message = "Hi"
	// Enable to see debug logs
	// bot.Debug = true

	u := chatbot.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			// log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			// log.Printf("[%s] %s", update.Message.From.UserName, update.Message.From.ID)
			// userId := update.Message.From.ID
			// userMsg := update.Message.Text


			bot.Send(chatbot.NewMessage(update.Message.Chat.ID, update.Message.Text))
		}
	}
}
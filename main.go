package main

import (
	"log"
	chatbot "main/bot"
	// tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	//TODO: store key
	bot, err := chatbot.NewBotAPI("6620296838:AAFGeDHI31nXwpxW0tObD0HwlUO9c1NyBro")
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
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			// msg := chatbot.NewMessage(update.Message.Chat.ID, update.Message.Text)
			// msg.ReplyToMessageID = update.Message.MessageID

			//TODO: remove ReplyToMessageID everywhere

			// bot.Send(msg)
			bot.Send(chatbot.NewMessage(update.Message.Chat.ID, update.Message.Text))
		}
	}
}
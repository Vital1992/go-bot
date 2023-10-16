package main

import (
	"log"
	chatbot "main/bot"
	"main/openai"
	"main/routes"
	"net/http"
	"os"
	// tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	router := routes.MovieRoutes()

	http.Handle("/api", router)

	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), router))

	bot, err := chatbot.NewBotAPI(os.Getenv("API_KEY"))
	if err != nil {
		log.Panic(err)
	}

	var useChatGPT = os.Getenv("USE_GPT")

	u := chatbot.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("Successfully received a message from user: %s", update.Message.From.UserName)

			var analyzedMsg = ""
			var messageToReturn = ""

			if useChatGPT == "true" {
				msgToAsk := "Return: 'positive' - user satisfied, 'concerns' - meaning user satisfied, but have some concerns and 'negative' - user are not satisfied. Return only one word."
				msgToAsk += " " + update.Message.Text
				log.Printf("Request to chat GPT: %s", msgToAsk)

				analyzedMsg = openai.SendToGPT(os.Getenv("GPT_KEY"), msgToAsk)
				log.Printf("Chat GPT response: %s", analyzedMsg)
			}

			switch analyzedMsg {
			case "positive":
				messageToReturn = "Thank you for your review! We're delighted to know that you're satisfied with your purchase!"
			case "concerns":
				messageToReturn = "Thank you for your review! If you have any questions or need assistance, please don't hesitate to contact our customer support team at 800-888-8888. We're here to help!"
			case "negative":
				messageToReturn = "We apologize that our product didn't meet your expectations. Please don't hesitate to get in touch with our customer support team at 800-888-8888 to discuss your concerns or to receive return instructions. We're here to assist you."
			default:
				messageToReturn = "Thank you for your review and for choosing us for your recent purchase! We greatly appreciate your support and trust in our products."
			}

			bot.Send(chatbot.NewMessage(update.Message.Chat.ID, messageToReturn))
		}
	}
}
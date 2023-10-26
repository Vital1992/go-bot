package db

import "main/chatbot/models"

var ProductList = []models.Product{

	{
		Id:          1,
		Name:        "Microwave",
		Description: "1.6 Cu. Stainless Stell",
		Brand:       "Panasonic",
		Category:    "Kitchen",
		Price:       204.99,
		Picture:     "https://pisces.bbystatic.com/image2/BestBuy_US/images/products/5834/5834501_sd.jpg",
	},
	{
		Id:          2,
		Name:        "Mixer",
		Description: "Artisan Series 5 Qt.",
		Brand:       "KitchenAid",
		Category:    "Kitchen",
		Price:       362.50,
		Picture:     "https://pisces.bbystatic.com/image2/BestBuy_US/images/products/6008/6008009_sd.jpg",
	},
	{
		Id:          3,
		Name:        "Monitor",
		Description: "27-inch 4K UHD",
		Brand:       "LG",
		Category:    "Computers",
		Price:       169.99,
		Picture:     "https://pisces.bbystatic.com/image2/BestBuy_US/images/products/6505/6505040_sd.jpg",
	},
}

// To keep track of conversation
var Conversation = []models.ChatCompletionMessage{}

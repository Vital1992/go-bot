package db

import "main/chatbot/models"

var ProductList = []models.Product{

	{
		Id:         	1,
		Name:      		"Microwave",
		Description:    "1.6 Cu. Stainless Stell",
		Brand:    		"Panasonic",
		Category:		"Kitchen",
	},
	{
		Id:         	2,
		Name:      		"Mixer",
		Description:    "600 Series 6 Qt.",
		Brand:    		"KitchenAid",
		Category:		"Kitchen",
	}, 
	{
		Id:         	3,
		Name:      		"Monitor",
		Description:    "27-inch 4K UHD",
		Brand:    		"LG",
		Category:		"Computers",
	},
}

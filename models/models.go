package models

type Product struct {
	Id        	 int 	`json:"id"`
	Name      	 string `json:"name"`
	Description  string `json:"description"`
	Brand      	 string `json:"brand"`
	Category 	 string `json:"category"`
}

type Message struct {
	Message string `json:"message"`
}
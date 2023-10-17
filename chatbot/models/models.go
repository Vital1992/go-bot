package models

type Product struct {
	Id        	 int 	`json:"id"`
	Name      	 string `json:"name"`
	Description  string `json:"description"`
	Brand      	 string `json:"brand"`
	Category 	 string `json:"category"`
	Price 	 	 float64 `json:"price"`
	Picture 	 string `json:"picture"`
}

type Message struct {
	Message string `json:"message"`
}

type AddProductRequest struct {
    Product Product `json:"product"`
}

type AddProductResponse struct {
    AddedProduct Product `json:"addedProduct"`
}

type GetAllProductsResponse struct {
    AllProducts []Product `json:"allProducts"`
}

type GetProductResponse struct {
    Product Product `json:"product"`
}

type DeleteProductResponse struct {
    DeletedProduct Product `json:"deletedProduct"`
	Message string `json:"message"`
}

type UpdateProductRequest struct {
    Product Product `json:"product"`
}

type UpdateProductResponse struct {
    UpdatedProduct Product `json:"updatedProduct"`
}

type ErrorResponse struct {
    Error   string `json:"error"`
}


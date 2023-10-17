package controller

import (
	"encoding/json"
	"log"

	"main/chatbot/db"
	models "main/chatbot/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddProduct(w http.ResponseWriter, r *http.Request) {
    var req models.AddProductRequest

    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&req.Product); err != nil {
        handleError(w, "Failed to decode request body", http.StatusBadRequest)
        return
    }

    for _, productToAdd := range db.ProductList {
        if productToAdd.Id == req.Product.Id {
            handleError(w, "Product with provided ID already exists", http.StatusNotAcceptable)
            return
        }
    }

    db.ProductList = append(db.ProductList, req.Product)

    // Prepare the response
    resp := models.AddProductResponse{AddedProduct: req.Product}

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    // Encode the response
    if err := json.NewEncoder(w).Encode(resp); err != nil {
        handleError(w, "Failed to encode response", http.StatusInternalServerError)
        return
    }
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	allProducts := db.ProductList
	resp := models.GetAllProductsResponse{AllProducts: allProducts}
	json.NewEncoder(w).Encode(resp)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
	   handleError(w, "Unable to convert string to int", http.StatusInternalServerError)
	   return
	}

	for _, product := range db.ProductList {

		if id == product.Id {
			resp := models.GetProductResponse{Product: product}
			json.NewEncoder(w).Encode(resp)
			return
		}
	}
	handleError(w, "Product not found", http.StatusNotFound)
}

func DeleteProductById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
	   handleError(w, "Unable to convert string to int", http.StatusInternalServerError)
	   return
	}

	for index, product := range db.ProductList {
		if id == product.Id {
			db.ProductList = append(db.ProductList[:index], db.ProductList[index+1:]...)
			resp := models.DeleteProductResponse{
				DeletedProduct: product,
				Message: "Thank yopu for the purchase of " + product.Brand + " " + product.Name,
			}
			json.NewEncoder(w).Encode(resp)
			return
		}
	}
	handleError(w, "Product not found", http.StatusNotFound)

}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var req models.UpdateProductRequest
	
	decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&req.Product); err != nil {
        handleError(w, "Failed to decode request body", http.StatusBadRequest)
        return
    }

	for index, productToUpdate := range db.ProductList {

		if productToUpdate.Id == req.Product.Id {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			db.ProductList[index] = req.Product
			resp := models.UpdateProductResponse{UpdatedProduct: db.ProductList[index]}
			json.NewEncoder(w).Encode(resp)
			return
		}
	}
	handleError(w, "No product found with provided ID", http.StatusNotFound)
}

func handleError(w http.ResponseWriter, message string, statusCode int) {
    response := models.ErrorResponse{Error: message}
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(statusCode)
    if err := json.NewEncoder(w).Encode(response); err != nil {
        log.Printf("Failed to send error response: %v", err)
    }
}
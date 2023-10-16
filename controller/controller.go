package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"main/db"
	"main/models"

	"github.com/gorilla/mux"
)

func AddProduct(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var product models.Product
	err := decoder.Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, productToAdd := range db.ProductList {
		if productToAdd.Id == product.Id {
			http.Error(w, "Product with provided ID already exists", http.StatusNotAcceptable)
			return
		}
	}

	db.ProductList = append(db.ProductList, product)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	allProducts := db.ProductList
	json.NewEncoder(w).Encode(allProducts)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
	   http.Error(w, "Unable to convert string to int", http.StatusInternalServerError)
	   return
	}

	for _, product := range db.ProductList {

		if id == product.Id {

			json.NewEncoder(w).Encode(product)
			return
		}
	}
	message := models.Message{Message: "Product Not found"}
	json.NewEncoder(w).Encode(message)
}

func DeleteProductById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
	   http.Error(w, "Unable to convert string to int", http.StatusInternalServerError)
	   return
	}

	for index, product := range db.ProductList {
		if id == product.Id {
			db.ProductList = append(db.ProductList[:index], db.ProductList[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(db.ProductList)

}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var product models.Product

	json.NewDecoder(r.Body).Decode(&product)

	for index, productToUpdate := range db.ProductList {

		if productToUpdate.Id == product.Id {

			db.ProductList[index] = product
			json.NewEncoder(w).Encode(product)
			return
		}
	}
	http.Error(w, "No product found with provided ID", http.StatusNotFound)
}
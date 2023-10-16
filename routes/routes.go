package routes

import (
	"encoding/json"
	"net/http"

	"main/controller"
	"main/models"

	"github.com/gorilla/mux"
)


func MovieRoutes() *mux.Router {
	var router = mux.NewRouter()
	router = mux.NewRouter().StrictSlash(true)

	//Home Toute
	router.HandleFunc("/api/",func(rw http.ResponseWriter, r *http.Request) {
		 message := models.Message{
			 Message: "Store API",
		 }
		json.NewEncoder(rw).Encode(message)
	})


	//Other Routes
	router.HandleFunc("/api/products",controller.AddProduct).Methods(http.MethodPost)
	router.HandleFunc("/api/products",controller.GetAllProducts).Methods(http.MethodGet)
	router.HandleFunc("/api/products/{id}",controller.GetProductById).Methods(http.MethodGet)
	router.HandleFunc("/api/products/{id}",controller.DeleteProductById).Methods(http.MethodDelete)
	router.HandleFunc("/api/products",controller.UpdateProduct).Methods(http.MethodPut)
	
	return router
}
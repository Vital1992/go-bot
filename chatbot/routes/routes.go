package routes

import (
	"encoding/json"
	"net/http"

	"main/chatbot/controller"
	"main/chatbot/models"

	"github.com/gorilla/mux"
)

func ProducRoutes() *mux.Router {
	// Initialize the router with StrictSlash
	router := mux.NewRouter().StrictSlash(true)

	// Home Route
	router.HandleFunc("/api/", func(rw http.ResponseWriter, r *http.Request) {
		message := models.Message{
			Message: "Store API",
		}
		json.NewEncoder(rw).Encode(message)
	})

	router.HandleFunc("/api/getMessage", controller.GetFromChat)

	// Product Routes
	router.HandleFunc("/api/products", controller.AddProduct).Methods(http.MethodPost)
	router.HandleFunc("/api/products", controller.GetAllProducts).Methods(http.MethodGet)
	router.HandleFunc("/api/products/{id}", controller.GetProductById).Methods(http.MethodGet)
	router.HandleFunc("/api/products/{id}", controller.DeleteProductById).Methods(http.MethodDelete)
	router.HandleFunc("/api/products", controller.UpdateProduct).Methods(http.MethodPut)

	// Apply CORS middleware
	router.Use(enableCORS)

	return router
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Continue with the next handler
		next.ServeHTTP(w, r)
	})
}

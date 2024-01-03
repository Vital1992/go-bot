package routes

import (
	"net/http"

	"main/chatbot/controller"

	"github.com/gorilla/mux"
)

func ProducRoutes() *mux.Router {
	// Initialize the router with StrictSlash
	router := mux.NewRouter().StrictSlash(true)

	// In programming, particularly in Go, you might encounter both json.NewEncoder(rw).Encode(message) and rw.Write(json.Marshal(message)) as ways to send JSON-encoded data to a writer (rw),
	// such as an HTTP response writer. While both achieve similar outcomes, they have some differences in their approach and usage:

	// json.NewEncoder(rw).Encode(message)
	// This line creates a new JSON encoder with json.NewEncoder, passing it the writer rw. It then calls Encode on the encoder to directly encode the message object into JSON and write it to rw.

	// err := json.NewEncoder(rw).Encode(message)
	// if err != nil {
	//     // handle error
	// }
	// Advantages:

	// Efficiency: It writes the JSON data directly to the writer without needing a separate buffer. This can be more efficient, especially for large objects.
	// Simplicity: It's a straightforward way to encode and write in one step.
	// Streaming: Ideal for streaming JSON data, as it doesn't require the entire JSON representation to be held in memory.
	// rw.Write(json.Marshal(message))
	// In this line, json.Marshal(message) is first called to encode the message object into JSON, which returns a byte slice and an error. This byte slice is then written to the writer rw using rw.Write.

	// data, err := json.Marshal(message)
	// if err != nil {
	//     // handle error
	// }
	// _, err = rw.Write(data)
	// if err != nil {
	//     // handle error
	// }
	// Advantages:

	// Flexibility: Allows you to manipulate the JSON byte slice before writing it to the writer.
	// Error Handling: Provides a clear separation of marshaling and writing steps, which can be useful for detailed error handling.
	// Disadvantages:

	// Memory Usage: Requires holding the entire JSON-encoded data in memory before writing it out, which might be less efficient for large data objects.
	// More Steps: Involves a separate marshaling step, adding a bit of complexity.
	// Conclusion
	// Use json.NewEncoder(rw).Encode(message) for direct, efficient writing of JSON data, especially when dealing with streaming or large objects.
	// Use rw.Write(json.Marshal(message)) if you need to handle or manipulate the JSON data before writing it or if you have a use case that requires separation of marshaling and writing steps.

	router.HandleFunc("/api/getMessage", controller.GetFromChat)

	// Product Routes
	// Home Route
	router.HandleFunc("/api/home", controller.GetHome).Methods(http.MethodGet, "OPTIONS")
	router.HandleFunc("/api/products", controller.AddProduct).Methods(http.MethodPost, "OPTIONS")
	router.HandleFunc("/api/products", controller.GetAllProducts).Methods(http.MethodGet, "OPTIONS")
	router.HandleFunc("/api/products/{id}", controller.GetProductById).Methods(http.MethodGet, "OPTIONS")
	router.HandleFunc("/api/products/{id}", controller.DeleteProductById).Methods(http.MethodDelete, "OPTIONS")
	router.HandleFunc("/api/products", controller.UpdateProduct).Methods(http.MethodPut, "OPTIONS")

	// Apply CORS middleware
	// router.Use(enableCORS)

	return router
}

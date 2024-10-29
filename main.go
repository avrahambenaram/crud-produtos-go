package main

import (
	"log"
	"net/http"

	"github.com/avrahambenaram/crud-produtos-go/internal/controller"
	"github.com/avrahambenaram/crud-produtos-go/internal/service"
)

func main() {
	server := http.NewServeMux()
	productService := &service.ProductService{}
	productController := controller.NewProductController(productService)

	server.Handle("/product/", http.StripPrefix("/product", productController.Mux))

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", server)
}

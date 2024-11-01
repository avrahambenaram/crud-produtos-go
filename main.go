package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/avrahambenaram/crud-produtos-go/internal/configuration"
	"github.com/avrahambenaram/crud-produtos-go/internal/controller"
	"github.com/avrahambenaram/crud-produtos-go/internal/service"
)

func main() {
	server := http.NewServeMux()
	productService := &service.ProductService{}
	productController := controller.NewProductController(productService)

	server.Handle("/product/", http.StripPrefix("/product", productController.Handler))

	log.Printf("Server running on port %d\n", configuration.Server.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", configuration.Server.Port), server)
}

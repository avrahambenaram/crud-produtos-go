package controller

import (
	"encoding/json"
	"net/http"

	"github.com/avrahambenaram/crud-produtos-go/internal/service"
)

type ProductController struct {
	Mux *http.ServeMux
	*service.ProductService
}

func NewProductController(productService *service.ProductService) *ProductController {
	mux := http.NewServeMux()
	productController := &ProductController{
		mux,
		productService,
	}

	mux.HandleFunc("GET /listall", productController.getAllProducts)

	return productController
}

func (c *ProductController) getAllProducts(w http.ResponseWriter, _ *http.Request) {
	products := c.ProductService.GetAllProducts()
	resp, err := json.Marshal(products)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(resp)
}

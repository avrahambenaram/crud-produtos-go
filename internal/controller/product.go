package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	mux.HandleFunc("GET /{ID}", productController.getProductByID)

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

func (c *ProductController) getProductByID(w http.ResponseWriter, r *http.Request) {
	pathID := r.PathValue("ID")
	id, err := strconv.Atoi(pathID)
	if err != nil {
		http.Error(w, "Insert a valid ID (non negative integer)", http.StatusForbidden)
		return
	}

	product, errService := c.ProductService.GetProductById(uint(id))
	if errService != nil {
		http.Error(w, errService.Error(), http.StatusNotFound)
		return
	}

	productJson, errJson := json.Marshal(product)
	if errJson != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(productJson)
}

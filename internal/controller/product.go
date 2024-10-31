package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/avrahambenaram/crud-produtos-go/internal/entity"
	"github.com/avrahambenaram/crud-produtos-go/internal/middleware"
	"github.com/avrahambenaram/crud-produtos-go/internal/service"
)

type ProductController struct {
	http.Handler
	*service.ProductService
}

func NewProductController(productService *service.ProductService) *ProductController {
	mux := http.NewServeMux()
	productController := &ProductController{
		middleware.SendJSON(mux),
		productService,
	}

	mux.HandleFunc("GET /listall", productController.getAllProducts)
	mux.HandleFunc("GET /{ID}", productController.getProductByID)
	mux.HandleFunc("POST /add", productController.addProduct)

	return productController
}

func (c *ProductController) getAllProducts(w http.ResponseWriter, r *http.Request) {
	products := c.ProductService.GetAllProducts()
	ctx := context.WithValue(r.Context(), "json", products)
	*r = *r.WithContext(ctx)
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
	fmt.Println("Got product", product.ID)

	ctx := context.WithValue(r.Context(), "json", product)
	*r = *r.WithContext(ctx)
	fmt.Println("sent context")
}

func (c *ProductController) addProduct(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "Content type must be json", http.StatusForbidden)
		return
	}

	product := entity.Product{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	productCreated, err := c.ProductService.InsertProduct(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	productJson, errJson := json.Marshal(productCreated)
	if errJson != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(productJson)
}

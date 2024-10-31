package controller

import (
	"context"
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
	mux.Handle(
		"POST /add",
		middleware.ParseBody(
			http.HandlerFunc(productController.addProduct),
		),
	)
	mux.Handle(
		"PUT /update/{ID}",
		middleware.ParseBody(
			http.HandlerFunc(productController.updateProduct),
		),
	)
	mux.HandleFunc("DELETE /delete/{ID}", productController.deleteProduct)

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

	ctx := context.WithValue(r.Context(), "json", product)
	*r = *r.WithContext(ctx)
}

func (c *ProductController) addProduct(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "Content type must be json", http.StatusForbidden)
		return
	}

	product := r.Context().Value("product").(entity.Product)
	productCreated, err := c.ProductService.InsertProduct(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctx := context.WithValue(r.Context(), "json", productCreated)
	*r = *r.WithContext(ctx)
}

func (c *ProductController) updateProduct(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "Content type must be json", http.StatusForbidden)
		return
	}

	pathID := r.PathValue("ID")
	id, err := strconv.Atoi(pathID)
	if err != nil {
		http.Error(w, "Insert a valid ID (non negative integer)", http.StatusForbidden)
		return
	}

	product := r.Context().Value("product").(entity.Product)
	product.ID = uint(id)
	productUpdated, errUpdated := c.ProductService.UpdateProduct(product)
	if errUpdated != nil {
		http.Error(w, errUpdated.Error(), http.StatusNotFound)
		return
	}

	ctx := context.WithValue(r.Context(), "json", productUpdated)
	*r = *r.WithContext(ctx)
}

func (c *ProductController) deleteProduct(w http.ResponseWriter, r *http.Request) {
	pathID := r.PathValue("ID")
	id, err := strconv.Atoi(pathID)
	if err != nil {
		http.Error(w, "Insert a valid ID (non negative integer)", http.StatusForbidden)
		return
	}

	errDel := c.ProductService.DeleteProduct(uint(id))
	if errDel != nil {
		http.Error(w, errDel.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(204)
}

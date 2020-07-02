package handlers

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/transport/http/jsonrpc"
	"github.com/gorilla/mux"
	"gokusyon/github.com/products-api/data"
	"log"
	"net/http"
	"strconv"
)

type Products struct {
	log *log.Logger
}

func NewProducts(log *log.Logger) *Products {
	return &Products{log}
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.log.Println("Handle GET Products")
	lp := dataimport.GetProducts()
	rw.Header().Add("Content-Type", jsonrpc.ContentType)
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to Encode Products", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.log.Println("Handle POST Products")
	prod := r.Context().Value(KeyProduct{}).(dataimport.Product)
	dataimport.AddProduct(&prod)
}

func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	p.log.Println("Handle Put Products")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err!= nil{
		http.Error(rw, "Unable to convert into number", http.StatusBadRequest)
	}

	prod := r.Context().Value(KeyProduct{}).(dataimport.Product)
	err = dataimport.UpdateProduct(id, &prod)
	if err == dataimport.ErrorProductNotFound{
		http.Error(rw, "Product not found" , http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Unable to update products" , http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct {}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := dataimport.Product{}
		err := prod.FromJson(r.Body)
		if err != nil{
			http.Error(rw, "", http.StatusBadRequest)
			return
		}

		// Validate the product
		err = prod.Validate()
		if err != nil{
			p.log.Println("[ERROR] validating product", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating Product: %s", err),
				http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}
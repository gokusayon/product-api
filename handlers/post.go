package handlers

import (
	dataimport "github.com/gokusayon/products-api/data"
	"net/http"
)
// swagger:route POST /products products addProduct
// Adds a product to the list of products
// responses:
//	200: productResponse
//	422: errorValidation
//	501: errorResponse

// AddProducts returns the products from datastore
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.log.Println("Handle POST Products")
	prod := r.Context().Value(KeyProduct{}).(dataimport.Product)
	dataimport.AddProduct(&prod)
}

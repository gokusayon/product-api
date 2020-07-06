package handlers

import (
	"net/http"

	dataimport "github.com/gokusayon/products-api/data"
)

// swagger:route POST /products products addProduct
// Adds a product to the list of products
// responses:
//	200: productResponse
//	422: errorValidation
//	501: errorResponse

// AddProduct returns the products from datastore
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.log.Debug("Handle POST Products")
	prod := r.Context().Value(KeyProduct{}).(dataimport.Product)
	p.productsDB.AddProduct(prod)
}

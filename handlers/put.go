package handlers

import (
	"github.com/gorilla/mux"
	dataimport "github.com/gokusayon/products-api/data"
	"net/http"
	"strconv"
)

// swagger:route PUT /products/{id} products updateProduct
// Updates a product
// responses:
//	201: noContent
//	422: errorValidation
//	501: errorResponse

// UpdateProducts updates a existing product in the datastore
func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	p.log.Println("Handle PUT Products")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err!= nil{
		http.Error(rw, "Unable to convert into number", http.StatusBadRequest)
	}
	p.log.Println("[DEBUG] updating record id", id)

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


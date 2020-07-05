package handlers

import (
	"net/http"

	dataimport "github.com/gokusayon/products-api/data"
)

// swagger:route PUT /products/{id} products updateProduct
// Updates a product
// responses:
//	201: noContent
//	422: errorValidation
//	501: errorResponse

// UpdateProducts updates a existing product in the datastore
func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	p.log.Debug("Handle PUT Products")

	id := getProductID(r)

	prod := r.Context().Value(KeyProduct{}).(dataimport.Product)
	err := p.productsDB.UpdateProduct(id, prod)

	if err == dataimport.ErrorProductNotFound {
		p.log.Error("Unable to find product", "id", id)

		rw.WriteHeader(http.StatusNotFound)
		dataimport.ToJSON(&GenericError{Message: "Product not found in database"}, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}

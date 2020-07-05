package handlers

import (
	dataimport "github.com/gokusayon/products-api/data"
	"net/http"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Deletes a Product
// responses:
//	201: noContent
//	501: errorResponse
//  501: errorResponse

// DeleteProducts deletes a product from datastore
func (p *Products) DeleteProducts(rw http.ResponseWriter, r *http.Request) {
	p.log.Println("Handle DELETE Products")

	id := getProductID(r)

	p.log.Println("[DEBUG] deleting record id", id)

	err := dataimport.DeleteProduct(id)
	if err == dataimport.ErrorProductNotFound{
		p.log.Println("[ERROR] deleting record id does not exist")

		rw.WriteHeader(http.StatusNotFound)
		dataimport.ToJson(&GenericError{Message: err.Error()}, rw)
		return
	}

	if err != nil {
		p.log.Println("[ERROR] deleting record", err)

		rw.WriteHeader(http.StatusInternalServerError)
		dataimport.ToJson(&GenericError{Message: err.Error()}, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
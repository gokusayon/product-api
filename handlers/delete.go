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
	p.log.Debug("Handle DELETE Products")

	id := getProductID(r)

	p.log.Debug("Deleting record", "id", id)

	err := p.productsDB.DeleteProduct(id)
	if err == dataimport.ErrorProductNotFound {
		p.log.Error("Unable to delete record. Id does not exist")

		rw.WriteHeader(http.StatusNotFound)
		dataimport.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	if err != nil {
		p.log.Error("Unable to delete record. Id does not exist")

		rw.WriteHeader(http.StatusInternalServerError)
		dataimport.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}

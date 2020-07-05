package handlers

import (
	// protos "github.com/gokusayon/currency/protos/currency"
	dataimport "github.com/gokusayon/products-api/data"
	"net/http"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
//	200: productsResponse

// GetProducts returns the products from datastore
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.log.Println("Handle GET Products")
	lp := dataimport.GetProducts()
	err := dataimport.ToJson(lp, rw)
	if err != nil {
		http.Error(rw, "Unable to Encode Products", http.StatusInternalServerError)
	}
}

// swagger:route GET /products/{id} products listSingleProduct
// Returns a product from list of products
// responses:
//	200: productResponse
//	404: errorResponse

// ListSingle returns a products based on ID from database
func (p *Products) ListSingle(rw http.ResponseWriter, r *http.Request) {
	p.log.Println("Handle GET Product By ID")
	id := getProductID(r)

	prod, err := dataimport.GetProductByID(id)

	switch err {
	case nil:

	case dataimport.ErrorProductNotFound:
		p.log.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusNotFound)
		dataimport.ToJson(&GenericError{Message: err.Error()}, rw)
		return
	default:
		p.log.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusInternalServerError)
		dataimport.ToJson(&GenericError{Message: err.Error()}, rw)
		return
	}

	// get exchange rate
	// rr := &protos.RateRequest{
	// 	Base: proto
	// }
	// p.cc.GetRate(context.Background())

	err = dataimport.ToJson(prod, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.log.Println("[ERROR] serializing product", err)
	}
}

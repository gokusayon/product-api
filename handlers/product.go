package handlers

import (
	protos "github.com/gokusayon/currency/protos/currency"
	dataimport "github.com/gokusayon/products-api/data"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// Products handler for getting and updating products
type Products struct {
	log      *log.Logger
	validate *dataimport.Validation
	cc       protos.CurrencyClient
}

// NewProducts returns a @Products handler
func NewProducts(log *log.Logger, v *dataimport.Validation, cc protos.CurrencyClient) *Products {
	return &Products{log, v, cc}
}

// A KeyProduct used for product object in context
type KeyProduct struct{}

// getProductID returns the product ID from the URL
// Panics if cannot convert the id into an integer
// this should never happen as the router ensures that
// this is a valid number
func getProductID(r *http.Request) int {
	// parse the product id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}

	return id
}

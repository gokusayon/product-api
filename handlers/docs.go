// Package classification for Product API
//
// Documentation for Product API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//  - application/json
//
//	Produces:
//  - application/json
// swagger:meta
package handlers

import dataimport "github.com/gokusayon/products-api/data"

// A list of products returns in response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []dataimport.Product
}

// A product returns in response
// swagger:response productResponse
type productResponseWrapper struct {
	// All products in the system
	// in: body
	Body dataimport.Product
}

// No content is returned by this API endpoint
// swagger:response noContent
type ProductsNoContent struct {}

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// swagger:parameters deleteProduct listSingleProduct updateProduct
type productIDParameterWrapper struct {
	// The ID of product to be deleted
	//	in: path
	//	required: true
	ID int `json:"id"`
}

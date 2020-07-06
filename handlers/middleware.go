package handlers

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/transport/http/jsonrpc"
	dataimport "github.com/gokusayon/products-api/data"
)

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		p.log.Debug("Inside Middleware")
		prod := dataimport.Product{}
		err := dataimport.FromJSON(&prod, r.Body)
		if err != nil {
			p.log.Error("Unable to deserialize product")
			rw.WriteHeader(http.StatusBadRequest)
			dataimport.ToJSON(&GenericError{Message: err.Error()}, rw)
			return
		}

		// Validate the product
		errs := p.validate.Validate(prod)
		if len(errs) != 0 {
			p.log.Error("Unable to validate product")

			// return the validation messages as an array
			rw.WriteHeader(http.StatusUnprocessableEntity)
			dataimport.ToJSON(&ValidationError{Messages: errs.Errors()}, rw)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		rw.Header().Add("Content-Type", jsonrpc.ContentType)
		next.ServeHTTP(rw, req)
	})
}

func (p *Products) MiddlewareContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", jsonrpc.ContentType)
		next.ServeHTTP(rw, r)
	})
}

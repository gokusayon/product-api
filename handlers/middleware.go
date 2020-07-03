package handlers

import (
	"context"
	"github.com/go-kit/kit/transport/http/jsonrpc"
	dataimport "gokusyon/github.com/products-api/data"
	"net/http"
)

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := dataimport.Product{}
		err := dataimport.FromJson(&prod, r.Body)
		if err != nil{
			p.log.Println("[ERROR] deserializing product", err)
			rw.WriteHeader(http.StatusBadRequest)
			dataimport.ToJson(&GenericError{Message: err.Error()}, rw)
			return
		}

		// Validate the product
		errs := p.validate.Validate(prod)
		if len(errs) != 0 {
			p.log.Println("[ERROR] validating product", errs)

			// return the validation messages as an array
			rw.WriteHeader(http.StatusUnprocessableEntity)
			dataimport.ToJson(&ValidationError{Messages: errs.Errors()}, rw)
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


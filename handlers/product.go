package handlers

import (
	"github.com/go-kit/kit/transport/http/jsonrpc"
	dataimport "gokusyon/github.com/products-api/data"
	"log"
	"net/http"
)

type Products struct {
	log *log.Logger
}

func NewProducts(log *log.Logger) *Products {
	return &Products{log}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}
	// Handle update

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := dataimport.GetProducts()
	rw.Header().Add("Content-Type", jsonrpc.ContentType)
	err := lp.TOJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to Marshal", http.StatusInternalServerError)
	}
}

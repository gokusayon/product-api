package handlers

import (
	"github.com/go-kit/kit/transport/http/jsonrpc"
	"gokusyon/github.com/products-api/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Products struct {
	log *log.Logger
}

func NewProducts(log *log.Logger) *Products {
	return &Products{log}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	// Handle request for a list of products
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// Handle a request to add a product
	if r.Method == http.MethodPost{
		p.addProduct(rw, r)
		return
	}

	// Handle an update request
	if r.Method == http.MethodPut {
		// Expect an ID in URI
		reg := regexp.MustCompile("/([0-9])+")
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			p.log.Println("More than one or no id")
			http.Error(rw , "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			p.log.Println("More than one group")
			http.Error(rw , "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			p.log.Println("Unable to convert to number", idString)
			http.Error(rw , "Invalid URI", http.StatusBadRequest)
			return
		}
		p.updateProducts(id, rw, r)
	}

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.log.Println("Handle GET Products")
	lp := dataimport.GetProducts()
	rw.Header().Add("Content-Type", jsonrpc.ContentType)
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to Encode Products", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.log.Println("Handle POST Products")

	prod := &dataimport.Product{}
	err := prod.FromJson(r.Body)

	if err != nil{
		http.Error(rw, "Unable to Decode Products", http.StatusInternalServerError)
	}

	dataimport.AddProduct(prod)
}

func (p *Products) updateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	p.log.Println("Handle Put Products")

	prod := &dataimport.Product{}
	err := prod.FromJson(r.Body)

	if err != nil{
		http.Error(rw, "Unable to Decode Products", http.StatusBadRequest)
		return
	}

	err = dataimport.UpdateProduct(id, prod)
	if err == dataimport.ErrorProductNotFound{
		http.Error(rw, "Product not found" , http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Unable to update products" , http.StatusInternalServerError)
		return
	}


}

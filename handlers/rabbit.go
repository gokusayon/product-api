package handlers

import (
	"net/http"

	dataimport "github.com/gokusayon/products-api/data"
	q "github.com/gokusayon/products-api/protos/product"
)

func (p *Products) PublishMessages(rw http.ResponseWriter, r *http.Request) {
	p.log.Info("Handling PublishMessages")
	prod := r.Context().Value(KeyProduct{}).(dataimport.Product)

	var msg q.ProductMessage
	msg = q.ProductMessage{
		Description: prod.Description,
		Name:        prod.Name,
		ID:          int32(prod.ID),
		Price:       prod.Price,
		SKU:         prod.SKU,
	}

	p.queue.Send(msg, p.queue.GetConfigUri())

}

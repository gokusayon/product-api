package handlers

import (
	"net/http"

	"github.com/gokusayon/mongo-microservice/data"
	protos "github.com/gokusayon/mongo-microservice/protos/product"
	queue "github.com/gokusayon/rabbitmq/queue"
	"github.com/gorilla/mux"
)

// CreateProduct updates the product. If not found it inserts the new product
func (p *Products) PublishMessages(rw http.ResponseWriter, r *http.Request) {
	p.log.Info("Handle PublishMessages")

	var prod protos.ProductMessage
	err := data.ToJSON(r.Body, &prod)

	if err != nil {
		p.log.Info("Error deserializing product", "err", err)
		data.ErrorWithJSON(rw, "Error deserializing product", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	q, ok := vars["q"]
	if !ok {
		p.log.Error("Unable to fetch queue name from request")
		data.ErrorWithJSON(rw, "Unable to fetch queue name from request", http.StatusBadRequest)
		return
	}

	msg := queue.MessageQueue{
		QueueName: q,
		Message:   &prod,
	}

	p.queue.Send(msg, q)
	if err != nil {
		p.log.Info("Error adding product to the db", "err", err)
		data.ErrorWithJSON(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.WriteHeader(http.StatusNoContent)
}

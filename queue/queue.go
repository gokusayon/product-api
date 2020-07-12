package queue

import (
	"math/rand"

	q "github.com/gokusayon/products-api/protos/product"
	hclog "github.com/hashicorp/go-hclog"
)

type ProductQueue struct {
	log          hclog.Logger
	config       Config
	pchan        chan QueueMessage
	subsciptions map[string](chan q.ProductMessage)
}

func NewProductQueue(log hclog.Logger, config Config) *ProductQueue {
	pq := &ProductQueue{
		log,
		config,
		make(chan QueueMessage, 50),
		map[string]chan q.ProductMessage{},
	}

	pq.handleQueue()

	return pq
}

func (q *ProductQueue) handleQueue() {
	q.log.Info("Starting consume/publish services", "uri", q.config.uri)
	go q.ConsumeMessages()
	go q.PublishMessages()
}

func (q *ProductQueue) GetConfigUri() string {
	return q.config.uri
}

func (pq *ProductQueue) Send(msg q.ProductMessage, queueName string) error {
	pq.log.Info("Sending message to queue")

	subs := QueueMessage{
		QueueName: queueName,
		Message:   msg,
	}

	uuid := Uid(10)

	var rchan = make(chan q.ProductMessage)
	pq.subsciptions[uuid] = rchan

	// msg := RabbitMsg{
	// 	QueueName: "gateway",
	// 	Message:   *docMsg,
	// }
	pq.pchan <- subs

	return nil

}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func Uid(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

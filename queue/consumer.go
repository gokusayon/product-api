package queue

import (
	"encoding/json"

	q "github.com/gokusayon/products-api/protos/product"

	"github.com/streadway/amqp"
)

func (pq *ProductQueue) ConsumeMessages() {

	pq.log.Info("Initialized ConsumeMessages go routine")

	conn, err := amqp.Dial(pq.config.uri)
	if err != nil {
		pq.log.Error("Unable to establish conn with rabbitmq", "err", err)
		panic(err)
	}
	defer conn.Close()

	cch, err := conn.Channel()
	if err != nil {
		pq.log.Error("Unable to create a channel", "err", err)
		panic(err)
	}
	defer cch.Close()

	// TODO: change this to database queue
	queue, err := cch.QueueDeclare(pq.config.gatewayQueue, false, false, false, false, nil)
	if err != nil {
		pq.log.Error("Failed to create queue", "err", err)
		panic(err)
	}

	deliveryChannel, err := cch.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		pq.log.Error("Failed to delivery channel", "err", err)
		panic(err)
	}

	for msg := range deliveryChannel {
		var data = q.ProductMessage{}
		err := json.Unmarshal(msg.Body, &data)
		if err != nil {
			pq.log.Error("Failed to deserialize data", "err", err)
			continue
		}
		pq.log.Info("Consuming message", "msg", data.String())

		err = msg.Ack(true)
		if err != nil {
			pq.log.Error("Failed to send acknowledgement", "err", err)
		}

		//  TODO: find waiting channel(with uid) and forward the reply to it
		// if rchan, ok := rchans[docRply.Uid]; ok {
		// 	rchan <- *docRply
		// }

	}

}

package queue

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

func (pq *ProductQueue) PublishMessages() {
	pq.log.Info("Initialized PublishMessages go routine")

	conn, err := amqp.Dial(pq.config.uri)
	if err != nil {
		pq.log.Error("Unable to establish conn with rabbitmq", "err", err)
		panic(err)
	}
	defer conn.Close()

	pc, err := conn.Channel()
	if err != nil {
		pq.log.Error("Unable to create a channel", "err", err)
		panic(err)
	}
	defer pc.Close()

	for msg := range pq.pchan {
		pq.log.Info("publishing message to the queue", "msg", msg.Message)

		data, err := json.Marshal(&msg.Message)
		if err != nil {
			pq.log.Error("Unable to serializing message", "err", err)
			continue
		}

		// TODO: change this to storage queue
		err = pc.Publish("", pq.config.gatewayQueue, false, false, amqp.Publishing{
			Body:        data,
			ContentType: "application/json",
		})
		if err != nil {
			pq.log.Error("Unable to publish message to queue", "err", err)
			continue
		}

		// wait for reply
	}
}

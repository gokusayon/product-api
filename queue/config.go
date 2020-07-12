package queue

import (
	"os"

	q "github.com/gokusayon/products-api/protos/product"
)

// product "github.com/gokusayon/products-api/protos/product"

type Config struct {
	uri           string
	databaseQueue string
	gatewayQueue  string
}

type QueueMessage struct {
	QueueName string
	Message   q.ProductMessage
}

type QueueReply struct {
	QueueName string
	Message   q.CreateProductReply
}

type Subsciptions struct {
	pchans map[string](chan q.ProductMessage)
}

func NewConfig() *Config {
	return &Config{
		uri:           getEnv("RABBIT_URI", "amqp://admin:admin@localhost:5672/"),
		gatewayQueue:  getEnv("RABBIT_GATEWAY_QUEUE", "gateway"),
		databaseQueue: getEnv("RABBIT_GATEWAY_QUEUE", "storage"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

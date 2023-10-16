package model

import (
	"github.com/go-redis/redis/v8"
	"github.com/rabbitmq/amqp091-go"
)

type MessageBroker struct {
	RedisClient     *redis.Client
	RabbitMqClient  *amqp091.Connection
	RabbitMqChannel *amqp091.Channel
}

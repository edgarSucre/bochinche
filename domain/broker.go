package domain

import amqp "github.com/rabbitmq/amqp091-go"

type QuoteMessage struct {
	Room    string `json:"room"`
	Message string `json:"message"`
}

type Broker interface {
	PublishQuoteRequest(QuoteMessage) error
	PublishResponseMessage(QuoteMessage) error
	GetQuoteConsummer() (<-chan amqp.Delivery, error)
	GetResponseConsummer() (<-chan amqp.Delivery, error)
}

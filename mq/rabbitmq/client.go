package rabbitmq

import (
	"context"
	"encoding/json"

	"github.com/edgarSucre/bochinche/domain"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQClient struct {
	ch            *amqp.Channel
	quoteQueue    *amqp.Queue
	responseQueue *amqp.Queue
}

func NewClient(channel *amqp.Channel) *RabbitMQClient {
	return &RabbitMQClient{ch: channel}
}

func (c *RabbitMQClient) Start() error {
	quote, err := c.ch.QueueDeclare(
		"quote",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	c.quoteQueue = &quote

	response, err := c.ch.QueueDeclare(
		"response",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	c.responseQueue = &response
	return nil
}

func (c *RabbitMQClient) PublishQuoteRequest(req domain.QuoteMessage) error {
	data, err := json.Marshal(req)
	if err != nil {
		return err
	}

	err = c.ch.PublishWithContext(
		context.Background(),
		"",
		c.quoteQueue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		},
	)

	return err
}

func (c *RabbitMQClient) PublishResponseMessage(response domain.QuoteMessage) error {
	data, err := json.Marshal(response)
	// if there is an error there is no way to deliver the resoponse
	if err != nil {
		return err
	}

	err = c.ch.PublishWithContext(
		context.Background(),
		"",
		c.responseQueue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(data),
		},
	)

	return err
}

func (c *RabbitMQClient) GetQuoteConsummer() (<-chan amqp.Delivery, error) {
	msgs, err := c.ch.Consume(
		c.quoteQueue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	return msgs, nil
}

func (c *RabbitMQClient) GetResponseConsummer() (<-chan amqp.Delivery, error) {
	msgs, err := c.ch.Consume(
		c.responseQueue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	return msgs, nil
}

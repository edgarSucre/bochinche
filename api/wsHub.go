package api

import (
	"encoding/json"
	"fmt"

	"github.com/edgarSucre/bochinche/domain"
	amqp "github.com/rabbitmq/amqp091-go"
)

type chatterRecord struct {
	room   string
	client *Client
}

type broadCastSignal struct {
	message []byte
	room    string
}

type Hub struct {
	clients       map[string]map[*Client]bool
	broadcast     chan broadCastSignal
	register      chan chatterRecord
	unregister    chan chatterRecord
	quoteResponse <-chan amqp.Delivery
	publishFn     func(domain.QuoteMessage) error
	repo          domain.ChatRepository
}

func newHub(repository domain.ChatRepository, qtResponse <-chan amqp.Delivery, fn func(domain.QuoteMessage) error) *Hub {

	return &Hub{
		broadcast:     make(chan broadCastSignal),
		register:      make(chan chatterRecord),
		unregister:    make(chan chatterRecord),
		quoteResponse: qtResponse,
		publishFn:     fn,
		clients:       make(map[string]map[*Client]bool),
		repo:          repository,
	}
}

func (h *Hub) run() {

	go func() {
		for d := range h.quoteResponse {
			var message domain.QuoteMessage
			err := json.Unmarshal(d.Body, &message)
			if err == nil {
				msg := fmt.Sprintf("Bot: %s", message.Message)
				h.broadcast <- broadCastSignal{message: []byte(msg), room: message.Room}
			}
		}
	}()

	for {
		select {
		case reg := <-h.register:
			if _, ok := h.clients[reg.room]; !ok {
				h.clients[reg.room] = make(map[*Client]bool)
			}
			h.clients[reg.room][reg.client] = true
		case reg := <-h.unregister:
			if _, ok := h.clients[reg.room]; ok {
				delete(h.clients[reg.room], reg.client)
				close(reg.client.send)
			}
		case signal := <-h.broadcast:
			for client := range h.clients[signal.room] {
				select {
				case client.send <- signal.message:
				default:
					close(client.send)
					delete(h.clients[signal.room], client)
				}
			}
		}
	}
}

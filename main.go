package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/edgarSucre/bochinche/api"
	"github.com/edgarSucre/bochinche/bot"
	"github.com/edgarSucre/bochinche/mq/rabbitmq"
	"github.com/edgarSucre/bochinche/repository/postgres"
	_ "github.com/jackc/pgx/v4/stdlib"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	// MQ setting
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("RabbitMQ: Could not connect")
	}
	defer connection.Close()

	channel, err := connection.Channel()
	if err != nil {
		log.Fatal("RabbitMQ could not create channel")
	}
	defer channel.Close()

	mqClient := rabbitmq.NewClient(channel)

	err = mqClient.Start()
	if err != nil {
		log.Fatal("RabbitMQ could not start")
	}

	bot := bot.New(mqClient)
	go bot.ListenForRequest()

	dbSource := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		"root",
		"secret",
		"localhost",
		"5432",
		"chat",
	)

	conn, err := sql.Open("pgx", dbSource)
	if err != nil {
		log.Fatal("Could not connect to the DB")
	}
	defer conn.Close()

	repository := postgres.NewRepository(conn)

	server := api.New(&repository)

	//Use logrus https://github.com/sirupsen/logrus
	log.Fatal(server.Start(mqClient))
}

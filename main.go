package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/edgarSucre/bochinche/api"
	"github.com/edgarSucre/bochinche/bot"
	"github.com/edgarSucre/bochinche/config"
	"github.com/edgarSucre/bochinche/mq/rabbitmq"
	"github.com/edgarSucre/bochinche/repository/postgres"
	_ "github.com/jackc/pgx/v4/stdlib"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	//environmnet reading
	env, err := config.GetEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	// MQ setting
	rabbitUrl := fmt.Sprintf(
		"amqp://%s:%s@%s:%s",
		env["RABBIT_USER"],
		env["RABBIT_PASS"],
		env["RABBIT_HOST"],
		env["RABBIT_PORT"],
	)
	connection, err := amqp.Dial(rabbitUrl)
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

	bot := bot.New(mqClient, env)
	go bot.ListenForRequest()

	dbSource := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		env["DB_USER"],
		env["DB_PASS"],
		env["DB_HOST"],
		env["DB_PORT"],
		env["DB_NAME"],
	)

	conn, err := sql.Open("pgx", dbSource)
	if err != nil {
		log.Fatal("Could not connect to the DB")
	}
	defer conn.Close()

	repository := postgres.NewRepository(conn)

	server := api.New(&repository)

	//Use logrus https://github.com/sirupsen/logrus
	log.Fatal(server.Start(mqClient, env))
}

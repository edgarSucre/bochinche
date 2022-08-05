package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/edgarSucre/bochinche/api"
	"github.com/edgarSucre/bochinche/repository/postgres"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {

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
	log.Fatal(server.Start())
}

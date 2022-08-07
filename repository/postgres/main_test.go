package postgres_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/edgarSucre/bochinche/config"
	db "github.com/edgarSucre/bochinche/repository/postgres"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

var (
	testDB *sql.DB
	repo   db.PostgresRepository
)

const (
	dbDriver = "pgx"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("config: could not load .env")
	}

	env, err := config.GetEnvironment()
	if err != nil {
		log.Fatal("Could not load environment: ", err)
	}

	dbSource := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		env["DB_USER"],
		env["DB_PASS"],
		env["DB_HOST"],
		env["DB_PORT"],
		os.Getenv("DB_TEST"),
	)

	testDB, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Could not connect to the DB", err)
	}

	repo = db.NewRepository(testDB)

	os.Exit(m.Run())
}

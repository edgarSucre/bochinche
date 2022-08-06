package config

import (
	"fmt"
	"os"
)

var env map[string]string

func GetEnvironment() (map[string]string, error) {
	if env != nil {
		return env, nil
	}

	env := make(map[string]string)

	env["DB_NAME"] = os.Getenv("DB_NAME")
	env["DB_PORT"] = os.Getenv("DB_PORT")
	env["DB_HOST"] = os.Getenv("DB_HOST")
	env["DB_USER"] = os.Getenv("DB_USER")
	env["DB_PASS"] = os.Getenv("DB_PASS")
	env["API_PORT"] = os.Getenv("API_PORT")
	env["SESSION_KEY"] = os.Getenv("SESSION_KEY")
	env["RABBIT_USER"] = os.Getenv("RABBIT_USER")
	env["RABBIT_PASS"] = os.Getenv("RABBIT_PASS")
	env["RABBIT_HOST"] = os.Getenv("RABBIT_HOST")
	env["RABBIT_PORT"] = os.Getenv("RABBIT_PORT")

	for k, v := range env {
		if v == "" {
			return nil, fmt.Errorf("config error: can't read %s, from environment", k)
		}
	}

	return env, nil

}

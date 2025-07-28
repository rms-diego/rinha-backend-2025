package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type env struct {
	PORT         string
	DATABASE_URL string
}

var Env *env

func NewConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	e := env{
		PORT:         os.Getenv("PORT"),
		DATABASE_URL: os.Getenv("DATABASE_URL"),
	}

	switch {
	case e.PORT == "":
		return fmt.Errorf("environment variable 'PORT' is not set")

	case e.DATABASE_URL == "":
		return fmt.Errorf("environment variable 'DATABASE_URL' is not set")

	default:
		Env = &e
		return nil
	}
}

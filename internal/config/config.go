package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type env struct {
	PORT                           string
	DATABASE_URL                   string
	PAYMENT_PROCESSOR_DEFAULT_URL  string
	PAYMENT_PROCESSOR_FALLBACK_URL string
}

var Env *env

func NewConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	switch {
	case os.Getenv("PORT") == "":
		return fmt.Errorf("environment variable 'PORT' is not set")

	case os.Getenv("DATABASE_USER") == "":
		return fmt.Errorf("environment variable 'DATABASE_USER' is not set")

	case os.Getenv("DATABASE_PASSWORD") == "":
		return fmt.Errorf("environment variable 'DATABASE_PASSWORD' is not set")

	case os.Getenv("DATABASE_HOST") == "":
		return fmt.Errorf("environment variable 'DATABASE_HOST' is not set")

	case os.Getenv("PAYMENT_PROCESSOR_DEFAULT_URL") == "":
		return fmt.Errorf("environment variable 'PAYMENT_PROCESSOR_DEFAULT_URL' is not set")

	case os.Getenv("PAYMENT_PROCESSOR_FALLBACK_URL") == "":
		return fmt.Errorf("environment variable 'PAYMENT_PROCESSOR_FALLBACK_URL' is not set")

	default:
		Env = &env{
			PORT: os.Getenv("PORT"),
			DATABASE_URL: fmt.Sprintf(
				"postgres://%v:%v@%v:5432/rinha_backend_2025",
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_PASSWORD"),
				os.Getenv("DATABASE_HOST"),
			),
			PAYMENT_PROCESSOR_DEFAULT_URL:  os.Getenv("PAYMENT_PROCESSOR_DEFAULT_URL"),
			PAYMENT_PROCESSOR_FALLBACK_URL: os.Getenv("PAYMENT_PROCESSOR_FALLBACK_URL"),
		}
		return nil
	}
}

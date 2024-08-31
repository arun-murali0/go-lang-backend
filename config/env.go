package config

import (
	"os"

	"github.com/joho/godotenv"
)

func DotEnv() error {
	goEnv := os.Getenv("GO_ENV")
	if goEnv == "" || goEnv == "development" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}
	return nil
}

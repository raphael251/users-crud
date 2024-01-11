package configs

import (
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type conf struct {
	DBDriver   string `env:"DB_DRIVER,required"`
	DBHost     string `env:"DB_HOST,required"`
	DBPort     string `env:"DB_PORT,required"`
	DBName     string `env:"DB_NAME,required"`
	DBUser     string `env:"DB_USER,required"`
	DBPassword string `env:"DB_PASSWORD,required"`

	ServerPort string `env:"SERVER_PORT,required"`
}

func LoadConfig() (*conf, error) {
	configs := conf{}

	// locally, APP_ENV is not present so it will load the envs from the .env file
	if os.Getenv("APP_ENV") == "" {
		godotenv.Load()
	}

	err := env.Parse(&configs)
	if err != nil {
		panic(err)
	}

	return &configs, nil
}

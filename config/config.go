package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
)

type Config struct {
	Env           string `env:"APP_ENV"`
	ServerPort    string `env:"APP_SERVER_PORT"`
	MongoHost     string `env:"MONGO_HOST"`
	MongoDatabase string `env:"MONGO_DATABASE"`
	MongoPort     string `env:"MONGO_PORT"`
	MongoUser     string `env:"MONGO_USER"`
	MongoPassword string `env:"MONGO_PASSWORD"`
}

func ParseConfig() (cfg *Config) {
	godotenv.Load(".env")
	switch os.Getenv("APP_ENV") {
	case "local":
		godotenv.Overload(".env")
		log.Println("Running on local environments")
		break
	default:
		log.Println("Running on production environments")
		break
	}

	cfg = new(Config)
	_, err := env.UnmarshalFromEnviron(cfg)
	if err != nil {
		fmt.Printf("Cannot read config %s\n", err.Error())
		os.Exit(1)
	}
	return
}

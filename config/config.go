package config

import (
	"os"
	"time"
)

type config struct {
	host     string
	database string
	port     string
	driver   string
	user     string
	password string
	timeout  time.Duration
}

// Config mongoDB
func newConfigMongoDB() *config {
	return &config{
		host:     os.Getenv("MONGO_HOST"),
		database: os.Getenv("MONGO_HOST"),
		password: os.Getenv("MONGO_ROOT_PASSWORD"),
		user:     os.Getenv("MONGO_ROOT_USER"),
		timeout:  10 * time.Second,
	}
}

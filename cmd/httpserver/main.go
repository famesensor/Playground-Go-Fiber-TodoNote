package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// todoRepository := mongo.NewMongoRepositotry()
	// todoService := todo.New()
	// todoHandler := handler.New()

	// fiber inti
	app := fiber.New()

	// server run
	log.Fatal(app.Listen(":30000"))
}

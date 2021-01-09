package protocol

import (
	"context"
	"log"

	"github.com/famesensor/playground-go-fiber-todonotes/config"
	"github.com/famesensor/playground-go-fiber-todonotes/internal/actors/database"
	interfaces "github.com/famesensor/playground-go-fiber-todonotes/internal/core/ports"
	"github.com/famesensor/playground-go-fiber-todonotes/internal/core/services"
	"github.com/famesensor/playground-go-fiber-todonotes/internal/handler"
	"github.com/famesensor/playground-go-fiber-todonotes/internal/repositories"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func ServerHttp() error {
	app := fiber.New()

	// Middleware
	app.Use(cors.New())
	app.Use(logger.New())

	// TODO: init validate

	// load config from .env
	cfg := config.ParseConfig()

	// Connect MongoDB
	MongoConn, err := database.NewMongoTodo(cfg)
	if err != nil {
		log.Fatal(err)
	}
	// Disconnect database
	defer func() {
		if err := MongoConn.TodoMongo.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	// init
	todoMongoRepo := repositories.NewMongoRepositotry(MongoConn.TodoMongo, MongoConn.TodoDB)
	todoService := services.NewTodoService(todoMongoRepo)
	todoHandler := handler.NewTodoHandler(todoService)

	// SetUp route
	setupRoute(app, todoHandler)

	log.Fatal(app.Listen(":" + cfg.ServerPort))
	return nil
}

func setupRoute(app *fiber.App, todoHandler interfaces.TodoHandler) {
	v1 := app.Group("/api/v1", logger.New())

	// Todo route
	todo := v1.Group("/todo")
	todo.Post("/create-todo", todoHandler.Create)
	todo.Patch("/update-todo", todoHandler.UpdateTodo)
	todo.Patch("/update-status", todoHandler.UpdateStatus)
	todo.Get("/todos", todoHandler.FindAll)
	todo.Get("/", todoHandler.FindById)
}

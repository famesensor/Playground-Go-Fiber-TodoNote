package protocol

import (
	"context"
	"log"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/famesensor/playground-go-fiber-todonotes/config"
	_ "github.com/famesensor/playground-go-fiber-todonotes/docs"
	"github.com/famesensor/playground-go-fiber-todonotes/internal/actors/database"
	interfaces "github.com/famesensor/playground-go-fiber-todonotes/internal/core/ports"
	"github.com/famesensor/playground-go-fiber-todonotes/internal/core/services"
	"github.com/famesensor/playground-go-fiber-todonotes/internal/handler"
	"github.com/famesensor/playground-go-fiber-todonotes/internal/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /api/v1
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
	app.Get("/docs/*", swagger.Handler) // default
	v1 := app.Group("/api/v1", logger.New())

	// Todo route
	todo := v1.Group("/todo")
	todo.Get("/todos", todoHandler.FindAll)
	todo.Post("/create-todo", todoHandler.Create)
	todo.Patch("/update-todo", todoHandler.UpdateTodo)
	todo.Patch("/update-status", todoHandler.UpdateStatus)
	todo.Get("/:id", todoHandler.FindById)
}

package main

import (
	"context"
	"os"

	//"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}
	app := fiber.New()
	client,err := mongo.Connect(context.TODO(),options.Client().ApplyURI("mongodb://localhost:27017/gomongo"))
	
	if err != nil{
			panic(err)
	}

	client.Database("gomongo").
	app.Use(cors.New())

	app.Listen(":" + port)
}

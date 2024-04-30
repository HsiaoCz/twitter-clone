package main

import (
	"log"

	"github.com/HsiaoCz/twitter-clone/twitter/conf"
	"github.com/HsiaoCz/twitter-clone/twitter/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/", handlers.HandleHome)

	app.Listen(conf.GetPort("PORT"))
}

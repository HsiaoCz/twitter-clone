package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/HsiaoCz/twitter-clone/twitter/conf"
	"github.com/HsiaoCz/twitter-clone/twitter/handlers"
	"github.com/HsiaoCz/twitter-clone/twitter/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(conf.GetMongoDBUrl("DBURL")))
	if err != nil {
		slog.Error("mongo db connect error", "err", err)
		return
	}

	var (
		port      = conf.GetPort("PORT")
		dbname    = conf.GetMongoDBName("DBNAME")
		userColl  = conf.GetMongoDBUserColl("USERCOLL")
		userStore = storage.NewMongoUserStore(client, dbname, userColl)
		store     = &storage.Store{
			User: userStore,
		}
		userHandler = handlers.NewUserHandler(store)
	)

	app := fiber.New()
	app.Use(compress.New())
	app.Use(logger.New(logger.ConfigDefault))

	app.Get("/", handlers.HandleHome)

	app.Get("/api/auth/signup", userHandler.HandleSignup)
	app.Get("/api/auth/login", userHandler.HandleLogin)
	app.Get("/api/auth/logout", userHandler.HandleLogout)

	go func() {
		if err := app.Listen(port); err != nil {
			panic(err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c

	now := time.Now()
	app.Shutdown()

	slog.Info("the server is shutting down", "since", time.Since(now))
}

package main

import (
	"context"
	"fmt"
	"github.com/saeedjhn/go-redis-event-listener/api/httpserver"
	"github.com/saeedjhn/go-redis-event-listener/configs"
	"github.com/saeedjhn/go-redis-event-listener/internal/bootstrap"
	"github.com/saeedjhn/go-redis-event-listener/internal/domain/event/userevent"
	"github.com/saeedjhn/go-redis-event-listener/internal/domain/listener/userlistener"
	"github.com/saeedjhn/go-redis-event-listener/internal/infrastructure/pubsub"
	"github.com/saeedjhn/go-redis-event-listener/pkg/cmd/migrations"
	"log"
	"os"
	"os/signal"
)

func main() {
	// Bootstrap
	app := bootstrap.App(configs.Development)
	log.Printf("%#v", app)

	// Migrations
	migrations.Up(app)

	// Start server
	server := httpserver.New(app)
	go func() {
		server.Serve()
	}()

	go func() {
		log.Println("Listen to running events")
		
		ctx := context.Background()
		ps := pubsub.New(app.RedisClient)

		// Define all listener
		userlistener.New(ps).Handler(ctx, userevent.QueuePattern)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt) // more SIGX (SIGINT, SIGTERM, etc)
	<-quit

	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, app.Config.Application.GracefulShutdownTimeout)
	defer cancel()

	if err := server.Router.Shutdown(ctxWithTimeout); err != nil {
		fmt.Println("http server shutdown error", err)
	}

	log.Println("received interrupt signal, shutting down gracefully..")
	// Close all db connection, etc
	app.CloseMysqlConnection()
	app.CloseRedisClientConnection()
	//app.ClosePostgresqlConnection() // Or etc..

	<-ctxWithTimeout.Done()
}

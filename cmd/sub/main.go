package main

// orderserver is a service that listens to nats-streaming-server
// and stores all incoming oreders (from the subject "orders")
// to the postgres database using in-memory cache

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/MrDavudov/wb_l0/nats_listener"
	"github.com/MrDavudov/wb_l0/server"
	"github.com/MrDavudov/wb_l0/cache"
	"github.com/MrDavudov/wb_l0/storage/postgres"
)

const (
	databaseURI = "postgresql://postgres:admin@localhost:5432/postgres"
	addr        = ":8080"
	clusterName = "wb-l0"
	clientID    = "orderServer"
	durableName = "orderSeverSub"
	subject     = "orders"
)

func main() {
	pg, err := postgres.NewStorage(databaseURI)
	if err != nil {
		log.Fatal(err)
	}
	defer logIfError(pg.Close)

	s, err := inmem.NewCache(inmem.WithPersistentStorage(pg))
	must(err)

	nl, err := nats_listener.New(clusterName, clientID, durableName, subject, s)
	must(err)
	defer logIfError(nl.Close)

	log.Println("NATS Listener started")

	server, err := server.New(addr, s)
	must(err)
	go logIfError(server.ListenAndServe)
	log.Printf("HTTP server is listening at %s", addr)
	defer func() {
		if err := server.Shutdown(context.Background()); err != nil {
			log.Println(err)
		}
	}()

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-sigint
	log.Println("Shutting down...")
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// logIfError is used to log an error thrown by the function that is running in a gouroutine
// or deferred function.
func logIfError(fn func() error) {
	if err := fn(); err != nil {
		log.Println(err)
	}
}

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/chrisjpalmer/korean-restaurants-be/internal/restdb"
	"github.com/chrisjpalmer/korean-restaurants-be/internal/server"
)

func main() {
	const connString = "postgres://postgres:postgres@localhost:5432/korean_restaurants"
	const port = "3001"

	// Init the database
	db, err := restdb.New(connString)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the server
	srv := server.New(db, port)

	log.Println("starting server on port", port)
	go func() {
		err := srv.Serve()
		if err != nil {
			log.Fatal("srv.Serve():", err)
		}
	}()

	// wait until we get the exit signal
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, syscall.SIGTERM, syscall.SIGINT)
	<-sigch

	log.Println("got exit signal... shutting down")

	err = srv.Close()
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}

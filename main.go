package main

import (
	"context"
	"gokusyon/github.com/products-api/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	log := log.New(os.Stdout, "products-api ", log.LstdFlags)

	helloHandler := handlers.NewHello(log)
	pingHandler := handlers.NewPing(log)
	productHandler := handlers.NewProducts(log)

	sm := http.NewServeMux()
	sm.Handle("/hello", helloHandler)
	sm.Handle("/ping", pingHandler)
	sm.Handle("/", productHandler)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		log.Println("Starting Server")
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sigChanel := make(chan os.Signal)
	signal.Notify(sigChanel, os.Kill)
	signal.Notify(sigChanel, os.Interrupt)

	sig := <-sigChanel
	log.Println("Recieved Signal for shutdown. Shutting down gracefully ...", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}

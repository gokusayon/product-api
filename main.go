package main

import (
	"context"
	"github.com/gorilla/mux"
	"gokusyon/github.com/products-api/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	log := log.New(os.Stdout, "products-api ", log.LstdFlags)

	pingHandler := handlers.NewPing(log)
	ph := handlers.NewProducts(log)

	sm := mux.NewRouter()
	sm.Handle("/ping", pingHandler)

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	putRouter.Use(ph.MiddlewareProductValidation)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(ph.MiddlewareProductValidation)

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

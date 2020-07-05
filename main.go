package main

import (
	"context"
	"github.com/go-openapi/runtime/middleware"
	protos "github.com/gokusayon/currency/protos/currency"
	dataimport "github.com/gokusayon/products-api/data"
	"github.com/gokusayon/products-api/handlers"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	log := log.New(os.Stdout, "products-api ", log.LstdFlags)
	v := dataimport.NewValidation()

	// Add grpc client
	conn, err := grpc.Dial("localhost:9092", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	cc := protos.NewCurrencyClient(conn)

	// Create the handlers
	ph := handlers.NewProducts(log, v, cc)

	// Create a new subrouter for add prefic and adding filter for response type
	router := mux.NewRouter()
	sm := router.PathPrefix("/products").Subrouter()
	sm.Use(ph.MiddlewareContentType)

	// Handle routes
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("", ph.GetProducts)
	getRouter.HandleFunc("/{id:[0-9]+}", ph.ListSingle)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	putRouter.Use(ph.MiddlewareProductValidation)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("", ph.AddProduct)
	postRouter.Use(ph.MiddlewareProductValidation)

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id:[0-9]+}", ph.DeleteProducts)

	ops := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(ops, nil)
	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

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

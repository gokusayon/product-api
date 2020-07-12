package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/go-openapi/runtime/middleware"
	protos "github.com/gokusayon/currency/protos/currency"
	data "github.com/gokusayon/products-api/data"
	"github.com/gokusayon/products-api/handlers"
	queue "github.com/gokusayon/products-api/queue"
	goHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

func getEnv() string {
	if env := os.Getenv("APP_ENV"); env != "" {
		return "currency:8081"
	} else {
		return "localhost:8081"
	}
}

func main() {

	log := hclog.Default()
	log.SetLevel(hclog.Trace)

	log.Info("Starting product-api server", "OS", runtime.GOOS)

	v := data.NewValidation()

	currency_host := getEnv()
	// Add grpc client
	log.Info("Initializing gRPC client", "host", currency_host)
	conn, err := grpc.Dial(currency_host, grpc.WithInsecure())

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	cc := protos.NewCurrencyClient(conn)
	productsDB := data.NewProductsDB(log, cc)

	config := queue.NewConfig()
	q := queue.NewProductQueue(log, *config)

	// Create the handlers
	ph := handlers.NewProducts(log, v, productsDB, q)

	// Create a new subrouter for add prefic and adding filter for response type
	router := mux.NewRouter()

	swaggerRouter := router.NewRoute().Subrouter()

	sm := swaggerRouter.PathPrefix("/products").Subrouter()
	sm.Use(ph.MiddlewareContentType)

	// Handle routes
	log.Info("Registering routes")
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("", ph.GetProducts).Queries("currency", "{[A-Z]{3}}")
	getRouter.HandleFunc("", ph.GetProducts)

	getRouter.HandleFunc("/{id:[0-9]+}", ph.ListSingle).Queries("currency", "{[A-Z]{3}}")
	getRouter.HandleFunc("/{id:[0-9]+}", ph.ListSingle)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	putRouter.Use(ph.MiddlewareProductValidation)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("", ph.AddProduct)
	postRouter.Use(ph.MiddlewareProductValidation)

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id:[0-9]+}", ph.DeleteProducts)

	log.Info("Registering swagger ..")
	ops := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(ops, nil)
	swaggerRouter.Handle("/docs", sh)
	swaggerRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// CORS
	ch := goHandlers.CORS(goHandlers.AllowedOrigins([]string{"*"}))

	// Publish messages to queue
	postRouter.HandleFunc("/publish", ph.PublishMessages)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      ch(router),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		ErrorLog:     log.StandardLogger(&hclog.StandardLoggerOptions{}), // set the logger for the server
	}

	go func() {
		log.Debug("Starting Server")
		err := s.ListenAndServe()
		if err != nil {
			log.Error("Unable to start server", "err", err)
		}
	}()

	sigChanel := make(chan os.Signal)
	signal.Notify(sigChanel, os.Kill)
	signal.Notify(sigChanel, os.Interrupt)

	sig := <-sigChanel
	log.Debug("Recieved Signal for shutdown. Shutting down gracefully ...", "sig", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}

package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber"

	"github.com/mackalex/building-microservices-youtube/product-api/data"
	"github.com/mackalex/building-microservices-youtube/product-api/handlers"
)

func main() {

	l := log.New(os.Stdout, "products-api ", log.LstdFlags)
	v := data.NewValidation()

	// create the handlers
	ph := handlers.NewProducts(l, v)

	// create a new serve mux and register the handlers
	// sm := mux.NewRouter()
	app := fiber.New()

	// handlers for API

	// GET /products
	app.Get("/products", ph.ListAll)

	// GET /products/1
	app.Get("/products/:id", ph.ListSingle)

	// putR := sm.Methods(http.MethodPut).Subrouter()
	// putR.HandleFunc("/products", ph.Update)
	// putR.Use(ph.MiddlewareValidateProduct)

	// postR := sm.Methods(http.MethodPost).Subrouter()
	// postR.HandleFunc("/products", ph.Create)
	// postR.Use(ph.MiddlewareValidateProduct)

	// deleteR := sm.Methods(http.MethodDelete).Subrouter()
	// deleteR.HandleFunc("/products/{id:[0-9]+}", ph.Delete)

	// handler for documentation
	// opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	// sh := middleware.Redoc(opts, nil)

	// getR.Handle("/docs", sh)
	// getR.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// // create a new server
	// s := http.Server{
	// 	Addr:         "localhost:9090",  // configure the bind address
	// 	Handler:      sm,                // set the default handler
	// 	ErrorLog:     l,                 // set the logger for the server
	// 	ReadTimeout:  5 * time.Second,   // max time to read request from the client
	// 	WriteTimeout: 10 * time.Second,  // max time to write response to the client
	// 	IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	// }

	// // start the server
	// go func() {
	// 	l.Println("Starting server on port 9090")

	// 	err := s.ListenAndServe()
	// 	if err != nil {
	// 		l.Printf("Error starting server: %s\n", err)
	// 		os.Exit(1)
	// 	}
	// }()

	// // trap sigterm or interupt and gracefully shutdown the server
	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt)
	// signal.Notify(c, os.Kill)

	// // Block until a signal is received.
	// sig := <-c
	// log.Println("Got signal:", sig)

	// // gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	// ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	// s.Shutdown(ctx)

	app.Listen(":3000")
}

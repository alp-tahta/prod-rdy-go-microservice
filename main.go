package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alp-tahta/prod-rdy-go-microservice/internal/handler"
	"github.com/alp-tahta/prod-rdy-go-microservice/internal/repository"
	"github.com/alp-tahta/prod-rdy-go-microservice/internal/service"
)

func main() {
	port := ":8080"

	// Dependency Injections
	couchbase := repository.NewCouchbase()
	service := service.NewService(couchbase)
	handler := handler.NewHandler(service)

	// Routes
	router := http.NewServeMux()
	router.HandleFunc("GET /health", handler.HealthChecker)

	// Routes - Products
	router.HandleFunc("POST /product", handler.CreateProduct)
	router.HandleFunc("GET /product/{id}", handler.GetProduct)
	router.HandleFunc("PUT /product/{id}", handler.UpdateProduct)
	router.HandleFunc("DELETE /product/{id}", handler.DeleteProduct)

	s := &http.Server{
		Addr:           port,
		Handler:        router,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Starting server on Port%s\n", port)
	log.Fatal(s.ListenAndServe())
}

package app

import (
	"blockchain-parser/config"
	"blockchain-parser/internal/parser"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func RunParser() error {
	c := InitializeContainer()

	c.Provide(ProvideHTTPRouter)

	return c.Invoke(startServer)
}

func ProvideHTTPRouter(
	parserHandler *parser.ParserHandler,
) *chi.Mux {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	}))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/block/current", parserHandler.GetCurrentBlock)
	r.Get("/transactions", parserHandler.GetTransactions)
	r.Post("/subscribe", parserHandler.Subscribe)
	r.Delete("/unsubscribe", parserHandler.Unsubscribe)

	return r
}

func startServer(router *chi.Mux) error {
	port := fmt.Sprintf(":%s", config.GlobalConfig.Port)
	log.Printf("Starting server on %s", port)
	return http.ListenAndServe(port, router)
}

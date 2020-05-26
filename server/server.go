package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	handlers "handler/http"
)

func addOrgRouter(orgHandler *handlers.OrgHandler) http.Handler {
	router := chi.NewRouter()

	
	router.Get("/", orgHandler.GetClientInfo)


	return router
}

func setupGracefulShutdown(server *http.Server, idleConnsClosed chan struct{}) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	signal.Notify(sigint, syscall.SIGTERM)
	<-sigint

	if err := server.Shutdown(context.Background()); err != nil {
		log.Printf("HTTP server Shutdown: %v", err)
	}

	close(idleConnsClosed)
}

// Run ..
func Run(serverPort string) {
	router := chi.NewRouter()
	router.Use(
		middleware.Recoverer,
		middleware.Logger,
	)

	orgHandler := handlers.NewOrgHandler()
	router.Route("/", func(rt chi.Router) {
		rt.Mount("/api/v1", addOrgRouter(orgHandler))
		//rt.Get("/ping", handlers.GetPingHandler)
	})

	server := &http.Server{
		Addr:           fmt.Sprintf(":%s", serverPort),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println(fmt.Sprintf("Running server on port %s...", serverPort))

	idleConnsClosed := make(chan struct{})
	go setupGracefulShutdown(server, idleConnsClosed)

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("Server closed: %v", err)
	}
	<-idleConnsClosed
}

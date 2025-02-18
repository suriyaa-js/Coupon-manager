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

	"github.com/gin-gonic/gin"
	"suriyaa.com/coupon-manager/server/props"
)

type server struct {
	host   string
	port   string
	router *gin.Engine
}

// NewServer creates a new server instance, also configures the API handlers
func NewServer(config *props.Config) *server {

	// Set Gin to release mode in production, this will disable debug logs
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	return &server{
		host:   config.Server.AppHost,
		port:   config.Server.AppPort,
		router: router,
	}
}

func (s *server) ConfigureAPI(config *props.Config) {
	configureAPI(config, s)
}

// Serve starts the server and listens for incoming requests
func (s *server) Serve() error {
	srv := &http.Server{
		Addr:           fmt.Sprintf("%s:%s", s.host, s.port),
		Handler:        s.router,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		IdleTimeout:    60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}

	}()

	fmt.Printf("Started server on: %s\n", fmt.Sprintf("http://%s:%s", s.host, s.port))
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println(context.Background(), "Server is shutting down ...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Server Shutdown: ", err)
		return err
	}

	fmt.Println(context.Background(), "Until next time ...")
	return nil

}

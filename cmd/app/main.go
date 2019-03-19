package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/afirth/fm-test/api"

	"github.com/caarlos0/env" // clean env parsing
	"github.com/gorilla/mux"  // because YareGNI
)

// Config contains access and setup configuration
type Config struct {
	Username    string        `env:"USERNAME,required"`
	Password    string        `env:"PASSWORD,required"`
	GracePeriod time.Duration `env:"GRACEPERIOD" envDefault:"5s"`
	HTTPAddr    string        `env:"HTTPADDR" envDefault:":80"`
}

// main starts an HTTP server with basic timeouts and graceful shutdown
// it also initializes a gbdx.Client which is authenticated
func main() {

	cfg := Config{}
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatalf("bad config: %+v", err)
	}

	client, err := api.NewClient(cfg.Username, cfg.Password)
	if err != nil {
		log.Fatalf("username and password set?: unable to create oauth2 http client: %+v", err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/healthz", api.HealthCheckHandler).Methods(http.MethodGet)
	r.HandleFunc("/search", client.SearchHandler).Methods(http.MethodPost)

	srv := &http.Server{
		Addr:           cfg.HTTPAddr,
		WriteTimeout:   time.Second * 15,
		ReadTimeout:    time.Second * 15,
		IdleTimeout:    time.Second * 60,
		MaxHeaderBytes: 1 << 20, // 1MB
		Handler:        r,       // our gorilla/mux
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Printf("OK: HTTP service listening on %s", cfg.HTTPAddr)
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("ListenAndServe: %s", err)
		}
	}()

	// Trap sigint and sigterm and initiate graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	// Block until deadline "GracePeriod" if there are connections
	sig := <-c
	log.Printf("Caught %v, shutting down in max %v",
		sig, cfg.GracePeriod)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.GracePeriod)
	defer cancel()

	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}

package app

import (
	"fmt"
	"net/http"
	"time"
)

func Run() error {
	// setup config
	cfg, err := LoadConfig()
	if err != nil {
		return err
	}

	// setup database
	_, err = InitializeDatabase(&cfg.Database)
	if err != nil {
		return err
	}

	// setup application router
	r := InitializeRouter()

	// setup Server
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Server.Address),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// run server
	fmt.Printf("🚀 Application running at %s:%d \n", cfg.Server.URL, cfg.Server.Address)
	err = server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil

}

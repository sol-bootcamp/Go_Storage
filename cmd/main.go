package main

import (
	"clases/cmd/server"
	"clases/internal/config"
	"fmt"
	"log"
)

func main() {
	// r := chi.NewRouter()

	cfg := config.LoadConfig()

	log.Println("Starting server on :" + cfg.ServerAddress)
	app := server.NewServerChi(cfg)
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}

}

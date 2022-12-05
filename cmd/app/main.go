package main

import (
	"log"

	"github.com/ilyasbulat/rest_api/internal/app"
	"github.com/ilyasbulat/rest_api/internal/config"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	app.Start(cfg)
}

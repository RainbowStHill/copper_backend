package main

import (
	"log"
	"path/filepath"

	"github.com/rainbowsthill/copper_backend/config"
	identity_server "github.com/rainbowsthill/copper_backend/service/id/server"
)

func main() {
	cf, err := filepath.Abs("./config_files/config.yaml")
	if err != nil {
		log.Fatalf("Config file not exist: %v", err)
	}

	err = config.AddConfigFile(cf)
	if err != nil {
		log.Fatalf("Cannot load config file %s: %v", cf, err)
	}

	port, err := config.GetBuiltInTypeConfig[int](cf, []string{"service", "port"})
	if err != nil {
		log.Fatalf("Cannot get configuration service.port: %v", err)
	}

	identity_server.Run(*port)
}

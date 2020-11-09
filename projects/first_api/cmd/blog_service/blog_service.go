package main

import (
	"log"

	"github.com/hhhieu/golang-training/first_api/pkg/database"
	"github.com/hhhieu/golang-training/first_api/pkg/system"

	"github.com/hhhieu/golang-training/first_api/internal/pkg/config"
)

const configFile = "/home/configs/blog_service.yml"

func main() {
	c, err := config.LoadConfig(system.IOUtil{}, configFile)
	if err != nil {
		log.Fatalf("Loading configuration failed: %v", err)
	}
	p, err := database.NewProvider(c.Database)
	if err != nil {
		log.Fatalf("Creating provider failed: %v", err)
	}
	err = p.Open()
	if err != nil {
		log.Fatalf("Connecting database failed: %v", err)
	}
}

package main

import (
	"context"
	"github.com/astronely/financial-helper_microservices/apiGateway/internal/app"
	"log"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("Error creating app: %v", err)
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("Error starting app: %v", err)
	}
}

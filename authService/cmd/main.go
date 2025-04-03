package main

import (
	"context"
	"github.com/astronely/financial-helper_microservices/authService/internal/app"
	"log"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("Error creating app: %s", err)
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("Error starting app: %s", err)
	}
}

package main

import (
	"context"
	"github.com/astronely/financial-helper_microservices/userService/internal/app"
	"log"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("Error initializing app: %s", err)
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("Error starting app: %s", err)
	}
}

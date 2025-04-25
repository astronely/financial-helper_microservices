package main

import (
	"context"
	"github.com/astronely/financial-helper_microservices/noteService/internal/app"
)

func main() {
	ctx := context.Background()

	a := app.NewApp(ctx)

	err := a.Run()
	if err != nil {
		panic("Failed to start application" + err.Error())
	}
}

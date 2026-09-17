package main

import (
	"context"
	"log"
	"shelter-platform/internal/app"
)

func main() {
	ctx := context.Background()

	if err := app.Run(ctx); err != nil {
		log.Fatal(err)
	}
}

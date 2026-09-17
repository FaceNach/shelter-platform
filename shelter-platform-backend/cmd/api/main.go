package main

import (
	"context"
	"fmt"
	"log"
	"shelter-platform/internal/app"
)

func main() {
	fmt.Println("Starting server")

	ctx := context.Background()

	if err := app.Run(ctx); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/evan3v4n/Go-lang--Microservice/application"
)

func main() {
	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app: ", err)
	}

	cancel()
}

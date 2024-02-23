package main

import (
	"context"
	"fmt"

	"github.com/evan3v4n/Go-lang--Microservice/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())

	if err != nil {
		fmt.Println("failed to start app: ", err)
	}
}

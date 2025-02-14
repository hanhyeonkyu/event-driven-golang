package main

import (
	"context"
	"log"

	"github.com/hanhyeonkyu/event-driven-golang/cmd/api/factory"
)

func main() {
	app, err := factory.NewApplication()
	if err != nil {
		log.Fatalf("Error creating application: %s", err)
	}
	factory.ResgisterRoutes(app)
	factory.RegisterConsumers(app)

	ctx := context.Background()

	err = app.StartConsumingQueues(ctx)
	if err != nil {
		log.Fatalf("Error consumer queues: %s", err)
	}
	defer app.DisconnectQueue(ctx)

	err = app.RunServer(ctx)
	if err != nil {
		log.Fatalf("Error running server: %s", err)
	}
}

package main

import (
	"context"
	"log"

	"github.com/bilou4/google-tasks-cli/cmd"
	"github.com/bilou4/google-tasks-cli/config"
	. "github.com/bilou4/google-tasks-cli/constants"
	"google.golang.org/api/option"
	"google.golang.org/api/tasks/v1"
)

func main() {
	var err error

	ctx := context.Background()
	client := config.GetClient()

	Srv, err = tasks.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve tasks Client %v", err)
	}

	cmd.Execute()
}

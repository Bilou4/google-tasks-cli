package main

import (
	"context"
	"log"
	"os"

	"github.com/bilou4/google-tasks-cli/cmd"
	"github.com/bilou4/google-tasks-cli/config"
	. "github.com/bilou4/google-tasks-cli/constants"
	"github.com/jedib0t/go-pretty/v6/table"
	"google.golang.org/api/option"
	"google.golang.org/api/tasks/v1"
)

func main() {
	var err error
	Table = table.NewWriter()
	Table.SetOutputMirror(os.Stdout)
	ctx := context.Background()
	client := config.GetClient()

	Srv, err = tasks.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve tasks Client %v", err)
	}

	cmd.Execute()
}

package cmd

import (
	"errors"
	"fmt"
	"log"

	"github.com/bilou4/google-tasks-cli/api"
	. "github.com/bilou4/google-tasks-cli/constants"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get <LIST-NAME>",
	Short: "Returns all tasks from a list",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a list name")
		}
		listsIds, err = api.GetLists()
		if err != nil {
			return err
		}
		if _, ok := listsIds[args[0]]; ok {
			return nil
		}
		return errors.New(fmt.Sprintf("'%s' listname does not exist", args[0]))
	},
	Run: func(cmd *cobra.Command, args []string) {
		tasksIds, err := api.GetTasks(listsIds[args[0]].Id)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Tasks from list %v :\n", args[0])
		Table.AppendHeader(table.Row{"ID", "Title", "Notes", "Due date"})
		for _, task := range tasksIds {
			Table.AppendRow(table.Row{fmt.Sprintf("%s", task.Id), fmt.Sprintf("%s", task.Title), fmt.Sprintf("%s", task.Notes), fmt.Sprintf("%s", task.Due)})
		}
		Table.Render()
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

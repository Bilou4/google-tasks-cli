package cmd

import (
	"errors"
	"fmt"
	"log"

	"github.com/bilou4/google-tasks-cli/api"
	"github.com/spf13/cobra"
)

var (
	listsIds map[string]string
	err      error
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
		for listname := range listsIds {
			if listname == args[0] {
				return nil
			}
		}
		return errors.New("This list name does not exist")
	},
	Run: func(cmd *cobra.Command, args []string) {
		tasksIds, err := api.GetTasks(listsIds[args[0]])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Tasks from list %v :\n", args[0])
		for task := range tasksIds {
			fmt.Printf("\t%s\n", task)
		}
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

package cmd

import (
	"errors"
	"log"

	"github.com/bilou4/google-tasks-cli/api"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <LIST-NAME> <TASK-NAME>",
	Short: "adds a task in an existing list",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("requires a list name and a task name")
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
		err = api.AddTask(listsIds[args[0]], args[1])
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Task added")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

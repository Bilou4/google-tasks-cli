package cmd

import (
	"errors"
	"fmt"
	"log"

	"github.com/bilou4/google-tasks-cli/api"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done <LIST-NAME> <TASK-NAME>",
	Short: "marks a task as done",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("requires a list name and a task name")
		}
		listsIds, err = api.GetLists()
		if err != nil {
			return err
		}

		if _, ok := listsIds[args[0]]; !ok {
			return errors.New(fmt.Sprintf("'%s' listname does not exist", args[0]))
		}

		tasksIds, err = api.GetTasks(listsIds[args[0]].Id)
		if _, ok := tasksIds[args[1]]; ok {
			return nil
		}
		return errors.New(fmt.Sprintf("'%s' taskname does not exist", args[1]))
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := api.TaskDone(listsIds[args[0]].Id, tasksIds[args[1]].Id)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Task marks as done")
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

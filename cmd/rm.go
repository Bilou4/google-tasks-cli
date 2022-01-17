package cmd

import (
	"errors"
	"log"

	"github.com/bilou4/google-tasks-cli/api"
	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm <LIST-NAME> <TASK-NAME>",
	Short: "removes a task from an existing list",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("requires a list name and a task name")
		}
		listsIds, err = api.GetLists()
		if err != nil {
			return err
		}
		// TODO: change the error according to the case
		for listname := range listsIds {
			if listname == args[0] {
				break
			}
		}
		tasksIds, err = api.GetTasks(listsIds[args[0]])
		for taskname := range tasksIds {
			if taskname == args[1] {
				return nil
			}
		}
		return errors.New("This list/task name does not exist")
	},
	Run: func(cmd *cobra.Command, args []string) {
		err = api.RemoveTask(listsIds[args[0]], tasksIds[args[1]])
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Task removed")
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

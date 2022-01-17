package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCompletedTasksCmd represents the deleteCompletedTasks command
var deleteCompletedTasksCmd = &cobra.Command{
	Use:   "deleteCompletedTasks <LIST-NAME>",
	Short: "deletes all completed tasks from an existing list",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a list name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deleteCompletedTasks called")
	},
}

func init() {
	rootCmd.AddCommand(deleteCompletedTasksCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCompletedTasksCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCompletedTasksCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

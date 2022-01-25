package cmd

import (
	"errors"
	"fmt"
	"log"
	"regexp"

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

		if _, ok := listsIds[args[0]]; !ok {
			return errors.New(fmt.Sprintf("'%s' listname does not exist", args[0]))
		}
		if dueDate := cmd.Flag("due-date").Value.String(); dueDate != "" {
			re := regexp.MustCompile("((19|20)\\d\\d)-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])")
			if !re.MatchString(dueDate) {
				return errors.New(fmt.Sprintf("Due date is not in the right format: expected YYYY-MM-DD, got, '%s'", dueDate))
			}
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		err = api.AddTask(listsIds[args[0]].Id, args[1], cmd.Flag("due-date").Value.String(), cmd.Flag("notes").Value.String())
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
	addCmd.Flags().StringP("due-date", "d", "", "Due date for the task. Format: YYYY-MM-DD)")
	addCmd.Flags().StringP("notes", "n", "", "A note to add to the note.")

}

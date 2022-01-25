package cmd

import (
	"errors"
	"fmt"
	"log"

	"github.com/bilou4/google-tasks-cli/api"
	"github.com/spf13/cobra"
)

// renameListCmd represents the renameList command
var renameListCmd = &cobra.Command{
	Use:   "renameList <OLD-LIST-NAME> <NEW-LIST-NAME>",
	Short: "renames an existing list",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("requires the old list name and a new list name")
		}
		listsIds, err = api.GetLists()
		if err != nil {
			return err
		}

		// checks the oldListName exists
		if _, ok := listsIds[args[0]]; ok {
			return nil
		}

		return errors.New(fmt.Sprintf("'%s' listname does not exist", args[0]))
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := api.RenameList(listsIds[args[0]].Id, args[1])
		if err != nil {
			log.Fatal(err)
		}
		log.Println("List has been renamed")
	},
}

func init() {
	rootCmd.AddCommand(renameListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// renameListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// renameListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

package cmd

import (
	"errors"
	"fmt"

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
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("renameList called")
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

package cmd

import (
	"errors"
	"log"

	"github.com/bilou4/google-tasks-cli/api"
	"github.com/spf13/cobra"
)

// addListCmd represents the addList command
var addListCmd = &cobra.Command{
	Use:   "addList <LIST-NAME>",
	Short: "creates a new list",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a list name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		err = api.AddList(args[0])
		if err != nil {
			log.Fatal(err)
		}
		log.Println("List added")
	},
}

func init() {
	rootCmd.AddCommand(addListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

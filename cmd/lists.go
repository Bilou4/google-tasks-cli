package cmd

import (
	"fmt"
	"log"

	"github.com/bilou4/google-tasks-cli/api"
	"github.com/spf13/cobra"
)

// listsCmd represents the lists command
var listsCmd = &cobra.Command{
	Use:   "lists",
	Short: "returns all existing lists",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		listId, err := api.GetLists()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Task Lists:")
		for list := range listId {
			fmt.Printf("\t%s\n", list)
		}
	},
}

func init() {
	rootCmd.AddCommand(listsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

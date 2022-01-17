package cmd

import (
	"errors"
	"log"

	"github.com/bilou4/google-tasks-cli/api"
	"github.com/spf13/cobra"
)

// rmListCmd represents the rmList command
var rmListCmd = &cobra.Command{
	Use:   "rmList <LIST-NAME>",
	Short: "removes a list and all its tasks",
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
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		err = api.RemoveList(listsIds[args[0]])
		if err != nil {
			log.Fatal(err)
		}
		log.Println("List removed")
	},
}

func init() {
	rootCmd.AddCommand(rmListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

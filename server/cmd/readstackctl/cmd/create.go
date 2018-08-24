package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"

	"github.com/jlevesy/readstack/server/api"
)

var (
	itemName string
	itemURL  string
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new entry",
	Run: func(cmd *cobra.Command, args []string) {
		conn, client, err := initClient()

		if err != nil {
			log.Fatal(err)
		}

		defer conn.Close()

		result, err := client.Create(
			context.Background(),
			&api.CreateRequest{
				Name: itemName,
				Url:  itemURL,
			},
		)

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Successfully created item %d with name %s \n", result.GetId(), result.GetName())
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&itemName, "name", "n", "", "Item name")
	createCmd.Flags().StringVarP(&itemURL, "url", "u", "", "Item URL")
	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("url")
}

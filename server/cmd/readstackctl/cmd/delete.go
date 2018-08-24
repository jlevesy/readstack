package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"

	"github.com/jlevesy/readstack/server/api"
)

var (
	itemID int64
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an item by ID",
	Run: func(cmd *cobra.Command, args []string) {
		conn, client, err := initClient()

		if err != nil {
			log.Fatal(err)
		}

		defer conn.Close()

		_, err = client.Delete(
			context.Background(),
			&api.DeleteRequest{
				Id: itemID,
			},
		)

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Successfully deleted Item with id %d", itemID)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().Int64VarP(&itemID, "id", "i", 0, "Item ID")
	createCmd.MarkFlagRequired("id")
}

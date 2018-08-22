package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"github.com/jlevesy/readstack/server/api"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Create a new entry",
	Run: func(cmd *cobra.Command, args []string) {
		conn, client, err := initClient()

		if err != nil {
			log.Fatal(err)
		}

		defer conn.Close()

		res, err := client.Index(context.Background(), &api.IndexRequest{})
		if err != nil {
			log.Fatal("could not call index endpoint", err)
		}

		renderItems(res)
	},
}

func renderItems(result *api.IndexResult) {
	w := tabwriter.NewWriter(os.Stdout, 5, 10, 3, ' ', 0)
	defer w.Flush()

	fmt.Fprintln(w, "ID\tName\tLink")
	fmt.Fprintln(w, "--\t----\t----")

	for _, item := range result.GetItems() {
		fmt.Fprintf(w, "%d\t%s\t%s\n", item.GetId(), item.GetName(), item.GetUrl())
	}
}

func init() {
	rootCmd.AddCommand(getCmd)
}

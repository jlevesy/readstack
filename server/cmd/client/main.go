package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"google.golang.org/grpc"

	"github.com/jlevesy/readstack/server/api"
)

func main() {
	backend := flag.String("b", "localhost:8080", "address of the backend")

	flag.Parse()

	conn, err := grpc.Dial(*backend, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to %s: %v", *backend, err)
	}

	defer conn.Close()

	client := api.NewItemClient(conn)

	res, err := client.Index(context.Background(), &api.IndexRequest{})
	if err != nil {
		log.Fatal("Could not call index endpoint", err)
	}

	renderItems(res)
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

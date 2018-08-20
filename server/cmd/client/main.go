package main

import (
	"context"
	"flag"
	"log"

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

	log.Println(res)
}

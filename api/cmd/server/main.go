package main

import (
	"log"
	"os"
	"time"

	"github.com/jlevesy/envconfig"
)

const (
	readstackAppName = "Readstack"

	envSeparator = "_"
)

type config struct {
	PostgresURL    string
	ListenHost     string
	ListenPort     int
	HandlerTimeout time.Duration
	WebAssetsPath  string
}

const (
	defaultPostgresURL     = "postgres://root:root@localhost:5672/readstack"
	defaultListenHost      = ""
	defaultListenPort      = 8080
	defaultHandlerTimemout = 200 * time.Millisecond
	defaultWebAssetsPath   = "dist/web"
)

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	logger.Println("Starting the server...")

	config := config{defaultPostgresURL, defaultListenHost, defaultListenPort, defaultHandlerTimemout, defaultWebAssetsPath}

	if err := envconfig.New(readstackAppName, envSeparator).Load(&config); err != nil {
		log.Fatal(err)
	}

	logger.Printf("Loaded config %v", config)
}

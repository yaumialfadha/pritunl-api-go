package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/yaumialfadha/pritunl-api-go"
)

func main() {
	// Initialize the Pritunl API client
	client, err := pritunl.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve the server ID from an environment variable
	server := os.Getenv("PRITUNL_DATA_SERVER")

	// Create a context for the request
	ctx := context.Background()

	// Start the specified server
	_, err = client.ServerStart(ctx, server)
	if err != nil {
		log.Fatal(err)
	}

	// Print a success message
	fmt.Println("Server Started Successfully")
}

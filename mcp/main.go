package main

import (
	"go-ai/mcp/tools"
	"go-ai/pkg/db"
	"go-ai/provider"
	"log"

	mcp "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
	"github.com/samber/do"
)

func main() {

	client, err := db.InitDB()

	if err != nil {
		log.Fatal(err)
	}

	injector := do.New()

	do.ProvideValue(injector, client)

	provider.Provide(injector)

	// Create a transport for the server
	serverTransport := stdio.NewStdioServerTransport()

	// Create a new server with the transport
	server := mcp.NewServer(serverTransport)

	tools.Register(injector, server)

	// Start the server
	if err := server.Serve(); err != nil {
		panic(err)
	}
	// Keep the server running
	select {}
}

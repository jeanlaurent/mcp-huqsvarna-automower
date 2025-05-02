package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	s := server.NewMCPServer(
		"Automower",
		"1.0.0",
	)

	// Add tool
	tool := mcp.NewTool("Automowers Status",
		mcp.WithDescription("Get status of my automowers"),
	)

	// Add tool handler
	s.AddTool(tool, automowerHandler)

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

func automowerHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	keys := HusqvarnaKeys{
		ClientID:     os.Getenv("HUQSVARNA_CLIENT_ID"),
		ClientSecret: os.Getenv("HUQSVARNA_CLIENT_SECRET"),
	}

	status, err := getMowerStatus(keys)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	jsonBytes, err := json.Marshal(status)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return mcp.NewToolResultText(string(jsonBytes)), nil
}

func debug() {
	keys := HusqvarnaKeys{
		ClientID:     os.Getenv("HUQSVARNA_CLIENT_ID"),
		ClientSecret: os.Getenv("HUQSVARNA_CLIENT_SECRET"),
	}

	status, err := getMowerStatus(keys)
	if err != nil {
		log.Fatal(err)
	}
	jsonBytes, err := json.Marshal(status)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonBytes))
}

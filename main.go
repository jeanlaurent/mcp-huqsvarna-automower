package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	s := server.NewMCPServer(
		"HusqvarnaAutomower",
		"1.0.0",
	)

	tool := mcp.NewTool("Husqvarna Automowers Status",
		mcp.WithDescription("Get status of my husqvarna automowers"),
	)

	s.AddTool(tool, automowerHandler)

	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

func automowerHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Setup keys from https://developer.husqvarnagroup.cloud
	keys := HusqvarnaKeys{
		ClientID:     os.Getenv("HUSQVARNA_CLIENT_ID"),
		ClientSecret: os.Getenv("HUSQVARNA_CLIENT_SECRET"),
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

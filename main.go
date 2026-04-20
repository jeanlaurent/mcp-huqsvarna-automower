package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func envOrDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func main() {
	transport := flag.String("transport", envOrDefault("TRANSPORT", "stdio"), "Transport type: stdio or http")
	port := flag.String("port", envOrDefault("PORT", "8080"), "HTTP port (only used with http transport)")
	flag.Parse()

	s := server.NewMCPServer(
		"HusqvarnaAutomower",
		"1.0.0",
	)

	tool := mcp.NewTool("Husqvarna Automowers Status",
		mcp.WithDescription("Get status of my husqvarna automowers"),
	)

	s.AddTool(tool, automowerHandler)

	switch *transport {
	case "http":
		addr := fmt.Sprintf(":%s", *port)
		httpServer := server.NewStreamableHTTPServer(s)
		log.Printf("Streamable HTTP server listening on %s/mcp", addr)
		if err := httpServer.Start(addr); err != nil {
			log.Fatalf("HTTP server error: %v", err)
		}
	case "stdio":
		if err := server.ServeStdio(s); err != nil {
			log.Fatalf("Stdio server error: %v", err)
		}
	default:
		log.Fatalf("Unknown transport: %s (expected 'stdio' or 'http')", *transport)
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

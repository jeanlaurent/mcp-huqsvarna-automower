package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func envOrDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// automowerInput is the (empty) input struct for the automower status tool.
// The official SDK requires a concrete input type even when there are no parameters.
type automowerInput struct{}

// automowerOutput wraps the mower status JSON string.
type automowerOutput struct {
	Status string `json:"status" jsonschema:"the JSON-encoded automower status"`
}

func automowerHandler(ctx context.Context, req *mcp.CallToolRequest, _ automowerInput) (*mcp.CallToolResult, automowerOutput, error) {
	keys := HusqvarnaKeys{
		ClientID:     os.Getenv("HUSQVARNA_CLIENT_ID"),
		ClientSecret: os.Getenv("HUSQVARNA_CLIENT_SECRET"),
	}

	status, err := getMowerStatus(keys)
	if err != nil {
		return &mcp.CallToolResult{IsError: true, Content: []mcp.Content{
			&mcp.TextContent{Text: err.Error()},
		}}, automowerOutput{}, nil
	}

	jsonBytes, err := json.Marshal(status)
	if err != nil {
		return &mcp.CallToolResult{IsError: true, Content: []mcp.Content{
			&mcp.TextContent{Text: err.Error()},
		}}, automowerOutput{}, nil
	}

	text := string(jsonBytes)
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: text},
		},
	}, automowerOutput{Status: text}, nil
}

func main() {
	transport := flag.String("transport", envOrDefault("TRANSPORT", "stdio"), "Transport type: stdio or http")
	port := flag.String("port", envOrDefault("PORT", "8080"), "HTTP port (only used with http transport)")
	flag.Parse()

	s := mcp.NewServer(&mcp.Implementation{
		Name:    "HusqvarnaAutomower",
		Version: "1.0.0",
	}, nil)

	mcp.AddTool(s, &mcp.Tool{
		Name:        "Husqvarna Automowers Status",
		Description: "Get status of my husqvarna automowers",
	}, automowerHandler)

	switch *transport {
	case "http":
		addr := fmt.Sprintf(":%s", *port)
		handler := mcp.NewStreamableHTTPHandler(func(_ *http.Request) *mcp.Server {
			return s
		}, nil)
		log.Printf("Streamable HTTP server listening on %s/mcp", addr)
		if err := http.ListenAndServe(addr, handler); err != nil {
			log.Fatalf("HTTP server error: %v", err)
		}
	case "stdio":
		if err := s.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
			log.Fatalf("Stdio server error: %v", err)
		}
	default:
		log.Fatalf("Unknown transport: %s (expected 'stdio' or 'http')", *transport)
	}
}

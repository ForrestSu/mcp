package main

import (
	"log"
	"net/http"

	"github.com/ForrestSu/mcp/config"
	"github.com/ForrestSu/mcp/internal/battery"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func runToolServer(url string) {
	// new MCP server.
	server := newMcpServer()
	registerTool(server)
	// create the streamable HTTP handler.
	handler := mcp.NewStreamableHTTPHandler(func(req *http.Request) *mcp.Server {
		return server
	}, nil)
	if err := http.ListenAndServe(url, handler); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func registerTool(s *mcp.Server) {
	battery.RegisterTool(s)
}

// create an MCP server.
func newMcpServer() *mcp.Server {
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "AI-helper",
		Version: "1.0.0",
	}, nil)
	// Add MCP-level logging middleware.
	server.AddReceivingMiddleware(config.CreateLoggingMiddleware())
	return server
}

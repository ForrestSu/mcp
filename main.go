package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/ForrestSu/mcp/config"
	"github.com/ForrestSu/mcp/internal/battery"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	var url string
	flag.StringVar(&url, "url", "localhost:8000", "host to connect to/listen on")
	flag.Parse()
	log.Printf("MCP server listening on %s", url)
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

func registerTool(s *mcp.Server) {
	battery.RegisterTool(s)
}

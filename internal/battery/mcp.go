package battery

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// RegisterTool register battery tool
func RegisterTool(s *mcp.Server) {
	handler := func(ctx context.Context, req *mcp.CallToolRequest, _ any) (*mcp.CallToolResult, any, error) {
		info, err := GetBatteryInfo()
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: info},
			},
		}, nil, nil
	}
	// register tool
	mcp.AddTool(s, &mcp.Tool{
		Name:        "battery",
		Description: "Query Local Machine's battery status",
	}, handler)
}

package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

var httpAddr = flag.String("http", ":8080", "HTTP address to serve on")

func main() {
	flag.Parse()

	opts := &mcp.ServerOptions{}
	server := mcp.NewServer(&mcp.Implementation{Name: "gather"}, opts)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "lookup-machine-by-ip",
		Description: "Lookup machine information by IP address",
	}, lookupMachineTool)

	if *httpAddr != "" {
		handler := mcp.NewStreamableHTTPHandler(func(*http.Request) *mcp.Server {
			return server
		}, nil)
		log.Printf("Starting HTTP server on %s", *httpAddr)
		if err := http.ListenAndServe(*httpAddr, handler); err != nil {
			log.Fatal(err)
		}
	}
}

type LookupArgs struct {
	IP string `json:"ip" jsonschema:"IP address to lookup"`
}

func lookupMachineTool(ctx context.Context, request *mcp.CallToolRequest, args LookupArgs) (*mcp.CallToolResult, any, error) {
	log.Printf("Received lookup request for IP: %s", args.IP)
	
	output, err := exec.Command("grep", "-w", args.IP, "data/machines").Output()
	text := "No machine found for IP: " + args.IP
	if err == nil {
		text = strings.TrimSpace(string(output))
	}
	
	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: text}},
	}, nil, nil
}
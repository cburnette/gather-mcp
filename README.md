# Gather MCP Server

A Model Context Protocol (MCP) server that provides machine lookup functionality by IP address. The server searches through a local machines database file and returns detailed information about machines matching the requested IP.

## Features

- **IP Lookup**: Find machine information by searching for IP addresses in the machines database
- **MCP Integration**: Works seamlessly with Claude Code and other MCP-compatible clients
- **HTTP Server**: Provides both MCP protocol and HTTP access

## Setup Instructions

### Prerequisites

- Go 1.19 or later
- Git

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/cburnette/gather-mcp.git
   cd gather-mcp
   ```

2. Download Go dependencies:

   ```bash
   go mod download
   ```

3. Build the server:

   ```bash
   go build -o gather-mcp
   ```

4. Run the server:
   ```bash
   ./gather-mcp
   ```

The server will start on port 8080 by default. You can specify a different port with the `-http` flag:

```bash
./gather-mcp -http :9090
```

### Data File

The server has a `data/machines` file containing machine information with IP addresses. Each line contains an IP address and associated machine details. IP address are between 10.0.1.1 and 10.0.1.100.

## Claude Code Configuration

To use this MCP server with Claude Code, add it to your Claude Code MCP configuration:

```bash
claude mcp add --transport http gather http://localhost:8080
```

### Available Tools

- `lookup-machine-by-ip`: Search for machine information by IP address

Example usage in Claude Code:

```
Can you lookup the machine with IP 192.168.1.100?
```

## Development

To modify the server, edit `main.go` and rebuild:

```bash
go build -o gather-mcp
```

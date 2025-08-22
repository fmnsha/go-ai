# Go-AI MCP Server Setup Guide

## Overview
Your go-ai MCP server is working correctly! The issue is that chat applications need to be configured to discover and connect to your MCP server.

## ✅ What's Working
- MCP server implementation ✓
- Board creation functionality ✓
- MongoDB integration ✓
- Tool registration ✓

## 🔧 Configuration Needed

### For Claude Desktop

1. **Find your Claude Desktop config directory:**
   - Windows: `%APPDATA%\Claude\`
   - Create if it doesn't exist

2. **Create or edit `claude_desktop_config.json`:**
```json
{
  "mcpServers": {
    "go-ai": {
      "command": "D:\\go-ai\\mcp\\main.exe",
      "args": [],
      "env": {
        "DB_HOST": "localhost",
        "DB_PORT": "27017"
      }
    }
  }
}
```

3. **Restart Claude Desktop**

### For Other MCP Clients

Use the provided `mcp-config.json` in your project root as a reference.

## 🎯 Available Tools

Your MCP server currently exposes:

1. **add board** - Creates a new board with:
   - `name` (string): Board identifier
   - `title` (string): Display title
   - `fields` (array): Board field names
   - `views` (array): Available view types

## 🧪 Testing Your Server

1. **Build the server:**
```bash
cd D:\go-ai\mcp
go build -o main.exe main.go
```

2. **Test manually:**
```bash
# The server uses stdio transport
echo '{"jsonrpc":"2.0","id":1,"method":"tools/list","params":{}}' | ./main.exe
```

## 🚨 Common Issues

1. **"No tools or prompts"** - MCP server not configured in client
2. **Connection fails** - Check file paths in config
3. **Tool calls fail** - Verify MongoDB is running

## 📝 Next Steps

1. Configure Claude Desktop with the config above
2. Restart Claude Desktop
3. Check Tools & Integrations - you should see "go-ai"
4. Test the "add board" tool

## 🔄 Development

To add more tools, edit `mcp/tools/board.tools.go` and uncomment workspace tools in `mcp/tools/tools.go`.

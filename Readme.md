# Husqvarna Automower MCP Server

A Model Context Protocol (MCP) server that provides access to Husqvarna connected automower API, allowing AI assistants to query information about your automower status.

This calls the husqvarna remote API. You need to create credentials at https://developer.husqvarnagroup.cloud

## Features

Returns full status from Husqvarna Automowers as specified in their [API](https://developer.husqvarnagroup.cloud/apis/automower-connect-api?tab=status%20description%20and%20error%20codes)

## Prerequisites

You need a `ClientID` and `ClientSecret` generated through the [Husqvarna developer portal](https://developer.husqvarnagroup.cloud)

1. Go to https://developer.husqvarnagroup.cloud
2. Sign up/sign in
3. Go to My Applications at https://developer.husqvarnagroup.cloud/applications
4. Click "Create App", enter a name, leave localhost, click create
5. You will receive an Application Key (ClientID) and Application Secret (ClientSecret)

## Available Tools

### Husqvarna Automowers Status

Get detailed information about all automowers.

**Parameters:**

None

## With Docker

The easiest way is with `docker`. Make sure [Docker Desktop](https://www.docker.com/products/docker-desktop/) is running, then run:
`docker build -t husqvarna-automower-mcp .`

Then in Claude Desktop or your favorite MCP Client

```
{
  "mcpServers": {
    "automower": {
      "command": "docker",
      "args": [
        "run",
        "-i",
        "--rm",
        "-e",
        "HUSQVARNA_CLIENT_ID",
        "-e",
        "HUSQVARNA_CLIENT_SECRET",
        "husqvarna-automower-mcp"
      ],
      "env": {  
        "HUSQVARNA_CLIENT_ID": "YourClientID",
        "HUSQVARNA_CLIENT_SECRET": "YoutClientSecret"
      }
    }
  }
}
```

To run with Streamable HTTP transport instead of stdio, add `-e TRANSPORT=http`, `-e PORT=8080`, and `-p 8080:8080` to the `docker run` arguments.

## With Docker Compose (Streamable HTTP + MCP Inspector)

This runs the automower server in HTTP mode alongside the [MCP Inspector](https://github.com/modelcontextprotocol/inspector) for easy testing.

1. Create a `.env` file with your credentials:

```
HUSQVARNA_CLIENT_ID=YourClientID
HUSQVARNA_CLIENT_SECRET=YourClientSecret
```

> ⚠️ **Do not commit `.env`** — it contains secrets. It is already listed in `.gitignore`.

2. Start both services:

```bash
docker compose up
```

3. Open the Inspector UI at [http://localhost:6274](http://localhost:6274) — it is pre-filled to connect to the automower service.

4. Click **Connect** to start the session.

5. Use the **Husqvarna Automowers Status** tool to query your mowers.


## With a golang environement, without Docker

If not using Docker, you will need a Go development environment then
`go build *.go -o husqvarna-automower-mcp`

Then in Claude Desktop or your favorite MCP Client:
```
{
  "mcpServers": {
    "automower": {
      "command": "husqvarna-automower-mcp",
      "env": {  
        "HUSQVARNA_CLIENT_ID": "YourClientID",
        "HUSQVARNA_CLIENT_SECRET": "YoutClientSecret"
      }
    }
  }
}
```



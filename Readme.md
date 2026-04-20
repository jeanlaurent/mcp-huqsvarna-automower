# Husqvarna Automower MCP Server

A Model Context Protocol (MCP) server that provides access to Husqvarna connected automower API, allowing AI assistants to query information about your automower status.

This calls the husqvarna remote API. You need to create credentials at https://developer.husqvarnagroup.cloud

## Features

The **Husqvarna Automowers Status** tool returns a full snapshot of every mower connected to your Husqvarna account. Concrete data returned per mower includes:

- **Status & activity**: current mode (e.g. `MAIN_AREA`), activity (e.g. `MOWING`, `CHARGING`, `PARKED_IN_CS`), state (e.g. `IN_OPERATION`), and any active error codes
- **Battery**: charge percentage (e.g. `87%`)
- **Schedule**: next planned start timestamp, calendar tasks with days of week and start/duration times
- **Last known GPS position**: latitude and longitude coordinates
- **Statistics**: total cutting time, total drive distance, number of charging cycles, blade usage time
- **Settings**: cutting height, headlight mode
- **System info**: mower name, model, serial number

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

Running the server in a Docker container is the recommended approach for Claude Desktop: it keeps your Husqvarna credentials isolated from the AI model process, and Claude calls it over stdio without any network setup required.

Make sure [Docker Desktop](https://www.docker.com/products/docker-desktop/) is running, then build the image:

```bash
docker build -t husqvarna-automower-mcp .
```

Then add it to your Claude Desktop (or any MCP client) configuration:

```json
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
        "HUSQVARNA_CLIENT_SECRET": "YourClientSecret"
      }
    }
  }
}
```

### Running in Streamable HTTP mode

To expose the server as a persistent HTTP endpoint (useful for multi-agent stacks, Home Assistant, or any client that connects over HTTP rather than spawning a subprocess):

```bash
docker run --rm \
  -e TRANSPORT=http \
  -e PORT=8080 \
  -e HUSQVARNA_CLIENT_ID=YourClientID \
  -e HUSQVARNA_CLIENT_SECRET=YourClientSecret \
  -p 8080:8080 \
  husqvarna-automower-mcp
```

The server will be available at `http://localhost:8080/mcp`.

## With Docker Compose (Streamable HTTP + MCP Inspector)

Streamable HTTP mode exposes the server as a persistent HTTP endpoint rather than a subprocess. This is suitable for multi-agent stacks, Home Assistant integrations, and any MCP client that can't launch a subprocess directly.

This Docker Compose setup runs the automower server in HTTP mode alongside the [MCP Inspector](https://github.com/modelcontextprotocol/inspector) for easy testing.

1. Build the container image:

```bash
docker compose build
```

2. Create a `.env` file with your credentials:

```bash
cat > .env << 'EOF'
HUSQVARNA_CLIENT_ID=YourClientID
HUSQVARNA_CLIENT_SECRET=YourClientSecret
EOF
```

> ⚠️ **Do not commit `.env`** — it contains secrets. It is already listed in `.gitignore`.

3. Start both services:

```bash
docker compose up
```

4. Open the Inspector UI at [http://localhost:6274](http://localhost:6274) — it is pre-filled to connect to the automower service.

   > **Note:** `DANGEROUSLY_OMIT_AUTH=true` and `ALLOWED_ORIGINS` are required for the MCP Inspector to work correctly when run inside Docker. `HOST=0.0.0.0` binds to all interfaces (needed for Docker port mapping), but causes the Inspector's proxy token URL to use `0.0.0.0` instead of `localhost`. Disabling auth avoids this mismatch. Always open the Inspector at `http://localhost:6274`, never `http://0.0.0.0:6274`.

5. Click **Connect** to start the session.

6. Use the **Husqvarna Automowers Status** tool to query your mowers.

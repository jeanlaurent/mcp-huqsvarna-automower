# Huqsvarna Automower MCP Server

A Model Context Protocol (MCP) server that provides access to Huqsvarna connected automowerAPI, allowing AI assistants to query information about your automower status.

This is calling a remote API, you have to create credentials at https://developer.husqvarnagroup.cloud

## Features

Return full status from an Huqsvarna Automowers as stated by their [API](https://developer.husqvarnagroup.cloud/apis/automower-connect-api?tab=status%20description%20and%20error%20codes)

## Prerequisite

You need a `ClientID` and `ClientSecret` as generated through the [husqvarna developer portal](https://developer.husqvarnagroup.cloud)

1. Go to https://developer.husqvarnagroup.cloud
2. Sign up/in
3. Go to myapplications https://developer.husqvarnagroup.cloud/applications
4. Click on create app, put a name, leave localhost, click on create
5. You should retreive application Key as ClientID and Application secret as ClientSecret

## Available Tools

### Huqsvarna_Automowers_Status

Get detailed information about all the automowers.

**Parameters:**

None
 

## Run it The easy way

The easiest is with `docker` just
`docker build -t husqvarna-automower .`

Then in claude

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
        "HUQSVARNA_CLIENT_ID",
        "-e",
        "HUQSVARNA_CLIENT_SECRET",
        "am"
      ],
      "env": {  
        "HUQSVARNA_CLIENT_ID": "YourClientID",
        "HUQSVARNA_CLIENT_SECRET": "YoutClientSecret"
      }
    }
  }
}
```


## Run it The hard way

If not then you need a golang development environement and 
`go build *.go -o husqvarna-automower`

Then in claude
```
{
  "mcpServers": {
    "automower": {
      "command": "husqvarna-automower",
      "env": {  
        "HUQSVARNA_CLIENT_ID": "YourClientID",
        "HUQSVARNA_CLIENT_SECRET": "YoutClientSecret"
      }
    }
  }
}
```



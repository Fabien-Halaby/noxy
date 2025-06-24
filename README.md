# Noxy [https://github.com/Fabien-Halaby/noxy](https://roadmap.sh/projects/caching-server)

A simple HTTP caching proxy server written in Go. This project forwards HTTP requests to an origin server, caches responses in memory, and serves cached responses for subsequent requests. It includes features to clear the cache via a CLI command or an HTTP endpoint.

## Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Project Structure](#project-structure)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Testing](#testing)

## Overview
The Caching Proxy acts as an intermediary between clients and an origin server (e.g., `http://example.com`). It caches HTTP responses in memory to improve performance by reducing requests to the origin. Cached responses include a custom header `X-Cache: HIT`, while fresh responses include `X-Cache: MISS`. The cache can be cleared using a CLI command (`--clear`) or an HTTP endpoint (`POST /clear`).

## Features
- **HTTP Proxy**: Forwards GET requests to the specified origin server.
- **In-Memory Cache**: Stores responses in memory for fast retrieval.
- **Cache Indicators**: Adds `X-Cache: HIT` for cached responses and `X-Cache: MISS` for fresh ones.
- **Cache Clearing**:
  - CLI: `./noxy --clear`
  - HTTP: `POST /clear`
- **MVC Architecture**: Organized into Model (cache logic), View (HTTP responses), and Controller (server and CLI handling).

## Project Structure
```bash
  noxy/
  ├── controller/
  │   ├── config.go
  │   └── server.go
  ├── model/
  │   └── cache.go
  ├── view/
  │   └── view.go
  ├── main.go
  ├── go.mod
  ├── README.md
  └── .gitignore
  ```

## Requirements
- **Go**: Version 1.16 or higher (`go version` to check).
- **Internet Connection**: For fetching responses from the origin server (e.g., `http://example.com`).
- **Terminal**: For running CLI commands.

## Installation
1. **Clone the Repository**:
   ```bash
   git clone git@github.com:Fabien-Halaby/noxy.git
   cd caching-proxy
   ```
2. **Initialize Go Module**:
   ```bash
   go mod init noxy
   ```
3. **Build the Project**:
   ```bash
   go build -o noxy
   ```

## Installation
1. **Run the Proxy Server**:
   ```bash
   ./noxy --port 3000 --origin http://example.com
   ```
    **Output**
   ```bash
   Starting caching proxy server on http://localhost:3000
   Forwarding requests to http://example.com .
   ```
2. **Clear the Cache**:
   ```bash
   ./noxy --clear
   ```
   **Output**
   ```bash
   Cache cleared successfully.
   ```

## Testing
1. **Test a request**:
   ```bash
   curl -v http://localhost:3000/test
   ```
   or
   Navigate to http://localhost:3000
2. **Verify log**:
   ```bash
   Received request: GET /test
   Cache Get: key=/test
   Cache Miss: key=/test
   Proxying request to: http://example.com
   Target URL: http://example.com/test
   Received response: Status=200, BodyLength=...
   Cache Set: key=/test, status=200, body_length=...
   Sending response: Status=200, FromCache=false
   ```


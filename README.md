# protosocket

Protobuf streaming over WebSocket implementation.

## Description

Server streams protobuf messages over WebSocket connection. Each message contains:

- Title: "My proto message"
- ID: Incremental counter (1, 2, 3, ...)
- Text: "Message number {id}"

## Running

### Backend

```bash
cd backend
go run cmd/main.go
```

Server starts on `:8000` and exposes WebSocket endpoint at `/ws`.

### Frontend

Serve the project root directory (to allow loading the proto file):

```bash
python3 -m http.server 8080
```

Then open `http://localhost:8080/frontend/` in your browser.

## Project Structure

- `backend/` - Go server with WebSocket handler
- `frontend/` - JavaScript client using protobuf.js
- `proto/` - Protobuf message definitions

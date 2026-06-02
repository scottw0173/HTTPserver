# HTTP Server

A RESTful HTTP server written in Go that serves as the backend for a Twitter-like microblogging application. Built on Go's standard `net/http` library with a PostgreSQL database managed via sqlc.

## Features

- User creation via REST API
- Post and retrieve "chirps" (messages capped at 140 characters)
- Profanity filtering on chirp content
- File server for static assets with hit tracking middleware
- Admin endpoints for metrics and environment-gated reset functionality
- PostgreSQL persistence via sqlc-generated type-safe queries

## Requirements

- Go 1.21+
- PostgreSQL
- [sqlc](https://sqlc.dev/) (if regenerating database queries)

## Setup

1. Clone the repository
```bash
git clone https://github.com/scottw0173/HTTPserver.git
cd HTTPserver
```

2. Copy `.env.example` to `.env` and configure your database URL
```bash
cp .env.example .env
```

Your `.env` should contain:
- DB_URL=postgres://username:password@localhost:5432/dbname
- PLATFORM=dev

3. Run the server
```bash
go run *.go
```

The server starts on port `8080`.

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/healthz` | Health check |
| POST | `/api/users` | Create a new user |
| POST | `/api/chirps` | Post a new chirp |
| GET | `/api/chirps` | List all chirps |
| GET | `/api/chirps/{id}` | Get a chirp by ID |
| GET | `/admin/metrics` | View file server hit count |
| POST | `/admin/reset` | Reset hits and users (dev only) |

## Notes

- The `/admin/reset` endpoint is restricted to `dev` environments via the `PLATFORM` environment variable
- Chirps exceeding 140 characters are rejected with a 400 error

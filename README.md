# HTTP Server

A RESTful HTTP server written in Go that serves as the backend for a Twitter-like microblogging application. Built on Go's standard `net/http` library with a PostgreSQL database managed via sqlc.

## Features

- User creation with bcrypt password hashing
- JWT-based authentication with short-lived access tokens and longer-lived refresh tokens
- Refresh token revocation for secure logout
- Protected endpoints requiring valid Bearer token authorization
- Post and retrieve "chirps" (messages capped at 140 characters)
- Profanity filtering on chirp content
- Chirpy Red subscription management via authenticated webhook
- File server for static assets with hit tracking middleware
- Admin endpoints for metrics and environment-gated reset functionality
- PostgreSQL persistence via sqlc-generated type-safe queries

## Requirements

- Go 1.21+
- PostgreSQL
- [sqlc](https://sqlc.dev/) (if regenerating database queries)
- [goose](https://github.com/pressly/goose) (if running migrations)

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
| POST | `/api/login` | Login and receive JWT tokens |
| POST | `/api/refresh` | Refresh access token |
| POST | `/api/revoke` | Revoke refresh token |
| POST | `/api/chirps` | Post a new chirp (auth required) |
| GET | `/api/chirps` | List all chirps |
| GET | `/api/chirps/{id}` | Get a chirp by ID |
| DELETE | `/api/chirps/{id}` | Delete a chirp (auth required) |
| PUT | `/api/users` | Update user email/password (auth required) |
| POST | `/api/polka/webhooks` | Polka webhook handler |
| GET | `/admin/metrics` | View file server hit count |
| POST | `/admin/reset` | Reset hits and users (dev only) |

## Notes

- The `/admin/reset` endpoint is restricted to `dev` environments via the `PLATFORM` environment variable
- Chirps exceeding 140 characters are rejected with a 400 error
- JWT access tokens are short-lived; use the refresh endpoint to obtain a new one without re-authenticating

# HTTPserver

A RESTful HTTP server and API built in Go, backed by PostgreSQL. Implements full CRUD operations, request validation, environment-gated admin controls, and request metrics tracking. Built as part of an ongoing series of projects to develop practical backend engineering skills.

## Stack

- **Language:** Go
- **Database:** PostgreSQL
- **Query layer:** [sqlc](https://sqlc.dev/) — type-safe Go generated from raw SQL
- **Router:** Go standard library `net/http` (no framework)
- **Environment:** `godotenv` for local config, `.env`-based DB connection

## Features

- RESTful API with explicit HTTP method routing (`GET`, `POST`, `DELETE`)
- PostgreSQL integration via `database/sql` and the `pq` driver
- Type-safe database queries generated with sqlc
- Intentional HTTP status code usage (200, 201, 400, 403, 404, 500)
- Middleware for tracking file server hit counts
- Environment-gated admin reset endpoint (dev only, returns 403 in other environments)
- Request body validation with structured error responses
- JSON request decoding and response encoding with clean DB-to-API type mapping
- UUIDs for resource identification (`google/uuid`)
- Configurable server timeouts and max header size

## API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/api/healthz` | Health check |
| `POST` | `/api/users` | Create a user |
| `POST` | `/api/chirps` | Create a chirp (validated, filtered) |
| `GET` | `/api/chirps` | List all chirps |
| `GET` | `/api/chirps/{id}` | Get a single chirp by UUID |
| `GET` | `/admin/metrics` | View file server hit count |
| `POST` | `/admin/reset` | Reset hit count and users (dev only) |

## Project Structure

```
HTTPserver/
├── main.go           # Server setup, route registration, DB connection
├── handlers.go       # HTTP handler functions
├── helpers.go        # respondWithJSON, respondWithError, filterChirp
├── models.go         # API-layer structs (chirp, user, request types)
├── sqlc.yaml         # sqlc config
├── sql/              # Raw SQL schema and queries
└── internal/
    └── database/     # sqlc-generated type-safe DB layer
```

## Setup

1. Create a PostgreSQL database and set the connection string in a `.env` file:
   ```
   DB_URL=postgres://user:password@localhost:5432/dbname?sslmode=disable
   PLATFORM=dev
   ```

2. Run database migrations from the `sql/` directory.

3. Build and run:
   ```bash
   go build -o HTTPserver && ./HTTPserver
   ```

Server starts on port `8080`.

## Status

Active development. Core user and chirp endpoints are complete and functional.

Roadmap:
- **Authentication** — password hashing with Argon2ID, stored securely in the database; login endpoint returning a session token
- **Authorization** — access control layer restricting endpoints by user identity and role (in design)
- **Webhooks** — inbound webhook support for external event integration (e.g. email notifications)

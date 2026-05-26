package main

import (
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"github.com/scottw0173/HTTPserver/internal/database"
)

// sqlc string:
// postgres://postgres:reteeks@localhost:5432/chirpy

// migrations:
// goose postgres "postgres://postgres:reteeks@localhost:5432/chirpy?sslmode=disable" up

type apiConfig struct {
	fileserverHits atomic.Int32
	dbQueries      *database.Queries
	platform       string
	serverSecret   string
	polkaKey       string
}

type chirpRequest struct {
	Body    string    `json:"body"`
	User_id uuid.UUID `json:"user_id"`
}

type user struct {
	ID            uuid.UUID `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Email         string    `json:"email"`
	Token         string    `json:"token"`
	Refresh_token string    `json:"refresh_token"`
	Is_chirpy_red bool      `json:"is_chirpy_red"`
}

type createuserRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type chirp struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
	User_id   uuid.UUID `json:"user_id"`
}

type subscriptionData struct {
	UserID uuid.UUID `json:"user_id"`
}

type subscriptionRequest struct {
	Event string           `json:"event"`
	Data  subscriptionData `json:"data"`
}

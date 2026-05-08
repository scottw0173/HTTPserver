package main

import (
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"github.com/scottw0173/HTTPserver/internal/database"
)

// sqlc string:
// postgres://postgres:reteeks@localhost:5432/chirpy

type apiConfig struct {
	fileserverHits atomic.Int32
	dbQueries      *database.Queries
	platform       string
	serverSecret   string
}

type chirpRequest struct {
	Body    string    `json:"body"`
	User_id uuid.UUID `json:"user_id"`
}

type user struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
}

type createuserRequest struct {
	Password  string `json:"password"`
	Email     string `json:"email"`
	ExpiresIn int32  `json:"expires_in"`
}

type chirp struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
	User_id   uuid.UUID `json:"user_id"`
}

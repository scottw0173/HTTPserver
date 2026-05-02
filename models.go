package main

import (
	"sync/atomic"

	"github.com/scottw0173/HTTPserver/internal/database"
)

// sqlc string:
// postgres://postgres:reteeks@localhost:5432/chirpy

type apiConfig struct {
	fileserverHits atomic.Int32
	dbQueries      *database.Queries
}

type chirpRequest struct {
	Body string `json:"body"`
}

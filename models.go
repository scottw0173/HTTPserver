package main

import (
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

type chirpRequest struct {
	Body string `json:"body"`
}

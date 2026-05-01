package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	cfg.fileserverHits.Store(0)
	w.WriteHeader(http.StatusOK)
}

func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	params := chirpRequest{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Error decoding parameters %s", err)
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error decoding parameters: %s", err))
	}

	if len(params.Body) > 140 {
		respondWithError(w, 400, "Chirp is too long")
		return
	}

	filteredChirp := filterChirp(params.Body)
	respondWithJSON(w, 200, map[string]string{
		"cleaned_body": filteredChirp,
	})
}

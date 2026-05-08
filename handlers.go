package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/scottw0173/HTTPserver/internal/auth"
	"github.com/scottw0173/HTTPserver/internal/database"
)

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev" {
		respondWithError(w, 403, "Forbidden")
		return
	}
	cfg.fileserverHits.Store(0)
	cfg.dbQueries.DeleteAllUsers(r.Context())
	w.WriteHeader(http.StatusOK)
}

func (cfg *apiConfig) handlerPostChirp(w http.ResponseWriter, r *http.Request) {
	tokenString, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "unauthorized")
		return
	}
	userID, err := auth.ValidateJWT(tokenString, cfg.serverSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	decoder := json.NewDecoder(r.Body)

	params := chirpRequest{}
	err = decoder.Decode(&params)
	if err != nil {
		log.Printf("Error decoding parameters %s", err)
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error decoding parameters: %s", err))
	}

	if len(params.Body) > 140 {
		respondWithError(w, 400, "Chirp is too long")
		return
	}

	filteredChirp := filterChirp(params.Body)

	newChirp := database.CreateChirpParams{
		Body:   filteredChirp,
		UserID: userID,
	}
	newPost, err := cfg.dbQueries.CreateChirp(r.Context(), newChirp)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("error adding new chirp to database: %s", err))
		return
	}
	respondWithJSON(w, http.StatusCreated, databaseChirptoChirp(newPost))
}

func (cfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	params := createuserRequest{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("error decoding user: %s", err))
		return
	}
	hashed_password, _ := auth.HashPassword(params.Password)
	userParams := database.CreateUserParams{
		Email:          params.Email,
		HashedPassword: hashed_password,
	}

	user, err := cfg.dbQueries.CreateUser(r.Context(), userParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error creating user: %s", err))
		return
	}
	respondWithJSON(w, http.StatusCreated, databaseUsertoUser(user))
}

func (cfg *apiConfig) handlerListChirps(w http.ResponseWriter, r *http.Request) {
	chirps, err := cfg.dbQueries.ReturnChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("error acquiring chirps: %s", err))
	}

	var jsonChirps []chirp
	for _, chirp := range chirps {
		jsonChirp := databaseChirptoChirp(chirp)
		jsonChirps = append(jsonChirps, jsonChirp)
	}
	respondWithJSON(w, http.StatusOK, jsonChirps)
}

func (cfg *apiConfig) handlerReturnChirp(w http.ResponseWriter, r *http.Request) {
	chirpID := r.PathValue("id")
	u, err := uuid.Parse(chirpID)
	if err != nil {
		log.Fatalf("failed to parse UUID: %v", err)
	}
	chirp, err := cfg.dbQueries.ReturnSingleChirp(r.Context(), u)
	if err != nil {
		respondWithError(w, 404, fmt.Sprintf("cannot find chirp: %s", err))
	}
	respondWithJSON(w, http.StatusOK, databaseChirptoChirp(chirp))
}

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	params := createuserRequest{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, fmt.Sprintln("Incorrect email or password"))
		return
	}
	if params.ExpiresIn == 0 || params.ExpiresIn > 3600 {
		params.ExpiresIn = 3600
	}
	dbUser, err := cfg.dbQueries.ReturnUser(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, fmt.Sprintln("Incorrect email or password"))
		return
	}

	OK, err := auth.CheckPasswordHash(params.Password, dbUser.HashedPassword)
	if !OK || err != nil {
		respondWithError(w, http.StatusUnauthorized, fmt.Sprintln("Incorrect email or password"))
		return
	}
	tokenExpiration := time.Duration(params.ExpiresIn) * time.Second
	userToken, err := auth.MakeJWT(dbUser.ID, cfg.serverSecret, tokenExpiration)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, fmt.Sprintln("error creating JWT token"))
	}
	respondWithJSON(w, http.StatusOK, user{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Email:     dbUser.Email,
		Token:     userToken,
	})
}

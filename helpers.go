package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/scottw0173/HTTPserver/internal/database"
)

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	w.Write([]byte(msg))
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(code)
	w.Write(data)
}

func filterChirp(msg string) string {
	banned_words := [3]string{"kerfuffle", "sharbert", "fornax"}
	words := strings.Split(msg, " ")
	for j, word := range words {
		for i := range banned_words {
			if strings.ToLower(word) == banned_words[i] {
				words[j] = "****"
			}
		}
	}
	cleaned_msg := strings.Join(words, " ")
	return cleaned_msg
}

func databaseUsertoUser(dbUser database.User) user {
	return user{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Email:     dbUser.Email,
	}
}

func databaseChirptoChirp(dbChirp database.Chirp) chirp {
	return chirp{
		ID:        dbChirp.ID,
		CreatedAt: dbChirp.CreatedAt,
		UpdatedAt: dbChirp.UpdatedAt,
		Body:      dbChirp.Body,
		User_id:   dbChirp.UserID,
	}
}

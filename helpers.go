package main

import (
	"encoding/json"
	"net/http"
	"strings"
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

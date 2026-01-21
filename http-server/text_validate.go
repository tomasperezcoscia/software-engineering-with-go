package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func handlerTextValidator(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type validValues struct {
		Valid bool `json:"valid"`
	}
	type cleanedBodyValues struct {
		CleanedBody string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not decode parameters", err)
		return
	}

	const maxTextLength = 140
	if len(params.Body) > maxTextLength {
		respondWithError(w, http.StatusBadRequest, "Text too long", nil)
		return
	}

	filteredBody := filterProfaneWordsInText(params.Body)

	if filteredBody == params.Body {
		respondWithJson(w, http.StatusOK, validValues{true})
	} else {
		respondWithJson(w, http.StatusOK, cleanedBodyValues{filteredBody})
	}
}

func filterProfaneWordsInText(text string) string {
	profaneWords := []string{"kerfuffle", "sharbert", "fornax"}
	wordsInText := strings.Fields(text)
	for i, word := range wordsInText {
		for _, profaneWord := range profaneWords {
			if strings.EqualFold(profaneWord, word) {
				wordsInText[i] = "****"
			}
		}
	}
	return strings.Join(wordsInText, " ")
}

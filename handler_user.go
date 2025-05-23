package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/appu900/webscraper/internal/database"
	model "github.com/appu900/webscraper/models"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithJSON(w, 400, fmt.Sprintf("Error parsing json"))
	}
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithJSON(w, 400, fmt.Sprintf("Could not create user:", err))
		return
	}
	respondWithJSON(w, 200, model.DatabaseUserToUser(user))
}

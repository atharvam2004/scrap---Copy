package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/atharvam2004/rss-go/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `name`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondError(w, 400, fmt.Sprintf("json parsinf %v", err))
		return
	}
	user,er:=apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if er != nil {
		respondError(w, 400, fmt.Sprintf("couldnt create no user %v", er))
		return
	}
	respondWithJson(w, 200, databaseUserToUser(user))
}

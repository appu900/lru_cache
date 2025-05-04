package model

import (
	"time"

	"github.com/appu900/webscraper/internal/database"
	"github.com/google/uuid"
)

type UserModel struct {
	ID        uuid.UUID `json:"id`
	CreatedAt time.Time `json:"created_at`
	UpdatedAt time.Time `json:"updated_at`
	Name      string    `json:"name"`
}

func DatabaseUserToUser(dbUser database.User) UserModel {
	return UserModel{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
	}
}

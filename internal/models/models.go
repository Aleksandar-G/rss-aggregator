package models

import (
	"time"

	"github.com/Aleksandar-G/rss-aggregator/internal/database"
)

type User struct {
	ID        interface{} `json:"id"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Name      string      `json:"name"`
}

func DatabaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
	}
}

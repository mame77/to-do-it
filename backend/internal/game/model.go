package game

import (
	"time"

	"github.com/mame77/to-do-it/backend/internal/model"
)

// Keep package-local request types but reuse the shared Game entity.
type Game = model.Game

type CreateGameRequest struct {
	Title       string    `json:"title" binding:"required"`
	Platform    string    `json:"platform"`
	Genre       string    `json:"genre"`
	ReleaseDate time.Time `json:"release_date"`
}

type UpdateGameRequest struct {
	Title       string    `json:"title"`
	Platform    string    `json:"platform"`
	Genre       string    `json:"genre"`
	ReleaseDate time.Time `json:"release_date"`
}

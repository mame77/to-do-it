package model

import "time"

// Game represents a game entity used across packages.
type Game struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Title       string    `json:"title"`
	Platform    string    `json:"platform"`
	Genre       string    `json:"genre"`
	ReleaseDate time.Time `json:"release_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// FixedEvent is a user-created fixed schedule.
type FixedEvent struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Title     string    `json:"title"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

// Schedule represents an auto-generated play schedule.
type Schedule struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	GameID    string    `json:"game_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Status    string    `json:"status"`
}

// Motivation holds user motivation/points state.
type Motivation struct {
	UserID string `json:"user_id"`
	Points int    `json:"points"`
	Rank   string `json:"rank"`
	Level  int    `json:"level"`
}

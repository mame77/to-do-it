package game

import "time"

// Game は、games テーブルのレコードを表す構造体です。
type Game struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"` // 本来は認証から取るが、一旦固定値やリクエストで指定
	Title       string    `json:"title" binding:"required"`
	Platform    string    `json:"platform"`
	Genre       string    `json:"genre"`
	Status      string    `json:"status"`  // unstarted, playing, completed
	ReleaseDate time.Time `json:"release_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateGameRequest は、ゲーム作成時のリクエストボディです。
type CreateGameRequest struct {
	// UserIDを削除（serviceで固定値を入れるため）
	Title       string    `json:"title" binding:"required"`
	Platform    string    `json:"platform"`
	Genre       string    `json:"genre"`
	Status      string    `json:"status"`
	ReleaseDate time.Time `json:"release_date"`
}

// UpdateGameRequest は、ゲーム更新時のリクエストボディです。
type UpdateGameRequest struct {
	Title       string    `json:"title"`
	Platform    string    `json:"platform"`
	Genre       string    `json:"genre"`
	Status      string    `json:"status"`
	ReleaseDate time.Time `json:"release_date"`
}
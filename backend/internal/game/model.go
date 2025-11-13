package game

import "time"

// Game はユーザーが所持するゲームの情報を保持します
type Game struct {
    ID        int       `json:"id"`
    UserID    int       `json:"user_id"`
    Title     string    `json:"title"`
    Platform  string    `json:"platform"`
    Status    string    `json:"status"` // 例: "積んでる", "プレイ中", "クリア"
    EstHours  int       `json:"estimated_hours"` // 予想プレイ時間
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
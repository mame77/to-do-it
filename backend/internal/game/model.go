package game

import "time"

// Status ... ゲームのプレイ状況
type Status string

const (
	[cite_start]StatusPending    Status = "未開始" // [cite: 19]
	[cite_start]StatusInProgress Status = "プレイ中" // [cite: 19]
	[cite_start]StatusCleared    Status = "クリア済み" // [cite: 19]
)

[cite_start]// Game ... ユーザーが所持するゲーム [cite: 164]
type Game struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	[cite_start]Title     string    `json:"title"`   // [cite: 166]
	[cite_start]Genre     string    `json:"genre"`   // [cite: 167]
	[cite_start]Status    Status    `json:"status"`  // [cite: 168]
	CreatedAt time.Time `json:"created_at"`
}
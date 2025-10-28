package calendar

import "time"

// FixedEvent (固定予定) [cite: 56-58, 81]
// ユーザーが手動で登録する、スケジュール自動生成時に考慮すべき予定（仕事、授業など）
type FixedEvent struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`    // どのユーザーの予定か
	Title     string    `json:"title"`      // "授業", "仕事" など
	StartTime time.Time `json:"start_time"` // 開始日時
	EndTime   time.Time `json:"end_time"`   // 終了日時
}

// Schedule (自動生成されたゲームスケジュール) [cite: 72, 96-101]
// 自動生成APIによって作られる「いつ、どのゲームをプレイするか」の予定
type Schedule struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	GameID    string    `json:"game_id"`    // プレイするゲームのID (gameパッケージと連携)
	StartTime time.Time `json:"start_time"` // プレイ開始予定時刻
	EndTime   time.Time `json:"end_time"`   // プレイ終了予定時刻
	Status    string    `json:"status"`     // "予定", "完了", "スキップ" [cite: 48, 73]
}

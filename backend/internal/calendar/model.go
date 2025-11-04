package calendar

import "time"

// ScheduleStatus ... スケジュールの進捗
type ScheduleStatus string

const (
	StatusScheduled ScheduleStatus = "予定"
	StatusCompleted ScheduleStatus = "完了"
	StatusSkipped   ScheduleStatus = "スキップ"
)

// FixedEvent (固定予定)
// ユーザーが手動で登録する、自動生成時に避けるべき時間
type FixedEvent struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	[cite_start]Title     string    `json:"title"`      // "仕事", "授業" など [cite: 28]
	StartTime time.Time `json:"start_time"` // 開始日時
	EndTime   time.Time `json:"end_time"`   // 終了日時
}

// Schedule (自動生成されたゲームスケジュール)
// 自動生成APIによって作られる「いつ、どのゲームをプレイするか」の予定
type Schedule struct {
	ID        string         `json:"id"`
	UserID    string         `json:"user_id"`
	GameID    string         `json:"game_id"`
	GameTitle string         `json:"game_title"` // (表示用にゲームタイトルも保持)
	[cite_start]StartTime time.Time      `json:"start_time"` // [cite: 32]
	EndTime   time.Time      `json:"end_time"`
	Status    ScheduleStatus `json:"status"` // "予定", "完了", "スキップ"
}
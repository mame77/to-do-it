package score

import "time"

// Motivation (継続状態)
// ユーザーの現在のポイントやランクなど、継続による蓄積状態を保持します。
//
type Motivation struct {
	UserID string `json:"user_id"`
	Points int    `json:"points"` // ポイント
	Rank   string `json:"rank"`   // "ブロンズ" など
	Level  int    `json:"level"`  // モチベーションゲージのレベル
}

// PlayResult (プレイ結果)
// 特定のスケジュールに対するプレイの成否を表します。
// これは継続促進ロジックへのインプットとして使われます。
//
type PlayResult struct {
	ScheduleID string `json:"schedule_id"` // どのスケジュールに対する結果か
	Result     string `json:"result"`      // "success" (成功) or "failure" (失敗/スキップ)
}

// PlaySession はゲームのプレイ記録を保持します
// 特定のプレイセッションの詳細なログ（時間、コメントなど）を記録します。
type PlaySession struct {
    ID        int       `json:"id"`
    UserID    int       `json:"user_id"`
    GameID    int       `json:"game_id"`
    Duration  int       `json:"duration_minutes"` // プレイ時間（分）
    Comment   string    `json:"comment"`
    Success   bool      `json:"success"` // 継続促進ロジックのための成功/失敗フラグ
    PlayedAt  time.Time `json:"played_at"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
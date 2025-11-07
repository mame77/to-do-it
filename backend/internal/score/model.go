package score

// Motivation (継続状態) [cite: 79, 145, 147]
// ユーザーの現在のポイントやランク
type Motivation struct {
	UserID string `json:"user_id"`
	Points int    `json:"points"` // ポイント
	Rank   string `json:"rank"`   // "ブロンズ" など
	Level  int    `json:"level"`  // モチベーションゲージのレベル
}

// PlayResult (プレイ結果) [cite: 76]
// プレイが「成功」したか「失敗」したか
type PlayResult struct {
	ScheduleID string `json:"schedule_id"` // どのスケジュールに対する結果か
	Result     string `json:"result"`      // "success" (成功) or "failure" (失敗/スキップ)
}
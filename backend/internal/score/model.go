package score

// Motivation (継続状態)
[cite_start]// ユーザーの現在のポイントやランク [cite: 36, 41]
type Motivation struct {
	UserID string `json:"user_id"`
	[cite_start]Points int    `json:"points"` // [cite: 55]
	[cite_start]Rank   string `json:"rank"`   // [cite: 43]
	[cite_start]Level  int    `json:"level"`  // [cite: 44]
}

// PlayResult (APIリクエスト用)
// プレイ結果を報告するためのリクエストボディ
type PlayResult struct {
	ScheduleID string `json:"schedule_id"`
	Result     string `json:"result"` // "completed" (完了) or "skipped" (失敗/スキップ)
}
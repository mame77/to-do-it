package calendar

// --- Core Data Models (Frontend types/index.ts に対応) ---

// PlayStatus はゲームのプレイ状態を示します。
type PlayStatus string

const (
	StatusUnstarted PlayStatus = "unstarted"
	StatusPlaying   PlayStatus = "playing"
	StatusCompleted PlayStatus = "completed"
)

// Game はユーザーが登録したゲームの情報を保持します。
type Game struct {
	ID      string     `json:"id"`
	Title   string     `json:"title"`
	Genre   string     `json:"genre,omitempty"` // GameGenreに対応
	Status  PlayStatus `json:"status"`
	AddedAt string     `json:"addedAt"` // ISO 8601形式の日時
}

// Schedule はゲームのプレイ予定を保持します。
type Schedule struct {
	ID        string `json:"id"`
	GameID    string `json:"gameId"`
	Date      string `json:"date"`      // YYYY-MM-DD形式
	StartTime string `json:"startTime"` // HH:MM形式
	EndTime   string `json:"endTime"`   // HH:MM形式
	Completed bool   `json:"completed"`
	Skipped   bool   `json:"skipped"`
}

// FixedEvent は定期的な予定や特定の日の固定された予定を保持します。
type FixedEvent struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	DayOfWeek     []int  `json:"dayOfWeek"` // 0-6 (日-土)
	StartTime     string `json:"startTime"`
	EndTime       string `json:"endTime"`
	IsRecurring   bool   `json:"isRecurring"`
	SpecificDate  string `json:"specificDate,omitempty"` // YYYY-MM-DD形式
}


// --- Request/Response Models (HandlerとService間で利用) ---

// GenerateScheduleRequest はスケジュール生成に必要なパラメータを定義します。
type GenerateScheduleRequest struct {
	StartDate      string `json:"startDate"` // YYYY-MM-DD
	DaysToSchedule int    `json:"daysToSchedule"`
}

// ScheduleActionRequest はスケジュールの完了・スキップに必要なパラメータを定義します。
type ScheduleActionRequest struct {
	ScheduleID string `json:"scheduleId"`
	Action     string `json:"action"` // "complete" or "skip"
}
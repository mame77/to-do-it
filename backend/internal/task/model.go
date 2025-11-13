package task

import "time"

// FixedEvent は変更できない固定の予定を保持します
type FixedEvent struct {
    ID        int       `json:"id"`
    UserID    int       `json:"user_id"`
    Title     string    `json:"title"`
    Type      string    `json:"type"` // 例: "Work", "Class", "Appointment"
    StartTime time.Time `json:"start_time"`
    EndTime   time.Time `json:"end_time"`
    IsRecurring bool    `json:"is_recurring"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
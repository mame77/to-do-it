package calendar

import (
	"database/sql"
	"time"
)

// Repository (インターフェース)
type Repository interface {
	// 固定予定 (FixedEvent)
	GetFixedEventsByUserID(userID string, start time.Time, end time.Time) ([]FixedEvent, error)
	CreateFixedEvent(event *FixedEvent) error
	// ... (UpdateFixedEvent, DeleteFixedEvent も必要) [cite: 83-84]

	// スケジュール (Schedule)
	GetSchedulesByUserID(userID string, start time.Time, end time.Time) ([]Schedule, error)
	CreateSchedules(schedules []Schedule) error
	UpdateScheduleStatus(scheduleID string, status string) error // [cite: 73]
}

// postgresRepository (実装)
type postgresRepository struct {
	db *sql.DB
}

// NewRepository ... DB接続を受け取り、リポジトリを初期化
func NewRepository(db *sql.DB) Repository {
	return &postgresRepository{db: db}
}

// --- 固定予定 (FixedEvent) の実装 ---

func (r *postgresRepository) GetFixedEventsByUserID(userID string, start time.Time, end time.Time) ([]FixedEvent, error) {
	// TODO: DBから固定予定を取得するSQLクエリを実装
	// SELECT * FROM fixed_events WHERE user_id = $1 AND start_time < $2 AND end_time > $3
	var events []FixedEvent
	// ... (db.QueryContext...)
	return events, nil
}

func (r *postgresRepository) CreateFixedEvent(event *FixedEvent) error {
	// TODO: DBに固定予定を保存するSQLクエリを実装 [cite: 81]
	// INSERT INTO fixed_events (...) VALUES (...)
	return nil
}

// --- スケジュール (Schedule) の実装 ---

func (r *postgresRepository) GetSchedulesByUserID(userID string, start time.Time, end time.Time) ([]Schedule, error) {
	// TODO: DBから生成済みスケジュールを取得するSQLクエリを実装 [cite: 72]
	// SELECT * FROM schedules WHERE user_id = $1 AND start_time < $2 AND end_time > $3
	var schedules []Schedule
	// ... (db.QueryContext...)
	return schedules, nil
}

func (r *postgresRepository) CreateSchedules(schedules []Schedule) error {
	// TODO: DBに複数のスケジュールを保存するSQLクエリを実装 (トランザクション推奨)
	// INSERT INTO schedules (...) VALUES (...)
	return nil
}

func (r *postgresRepository) UpdateScheduleStatus(scheduleID string, status string) error {
	// TODO: DBのスケジュールステータスを更新するSQLクエリを実装 [cite: 73]
	// UPDATE schedules SET status = $1 WHERE id = $2
	return nil
}

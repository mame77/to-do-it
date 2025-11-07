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
	query := `SELECT id, user_id, title, start_time, end_time
			  FROM fixed_events
			  WHERE user_id = ? AND start_time < ? AND end_time > ?
			  ORDER BY start_time`

	rows, err := r.db.Query(query, userID, end, start)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []FixedEvent
	for rows.Next() {
		var event FixedEvent
		if err := rows.Scan(&event.ID, &event.UserID, &event.Title, &event.StartTime, &event.EndTime); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

func (r *postgresRepository) CreateFixedEvent(event *FixedEvent) error {
	query := `INSERT INTO fixed_events (id, user_id, title, start_time, end_time)
			  VALUES (?, ?, ?, ?, ?)`

	_, err := r.db.Exec(query, event.ID, event.UserID, event.Title, event.StartTime, event.EndTime)
	return err
}

// --- スケジュール (Schedule) の実装 ---

func (r *postgresRepository) GetSchedulesByUserID(userID string, start time.Time, end time.Time) ([]Schedule, error) {
	query := `SELECT id, user_id, game_id, start_time, end_time, status
			  FROM schedules
			  WHERE user_id = ? AND start_time < ? AND end_time > ?
			  ORDER BY start_time`

	rows, err := r.db.Query(query, userID, end, start)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []Schedule
	for rows.Next() {
		var schedule Schedule
		if err := rows.Scan(&schedule.ID, &schedule.UserID, &schedule.GameID, &schedule.StartTime, &schedule.EndTime, &schedule.Status); err != nil {
			return nil, err
		}
		schedules = append(schedules, schedule)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return schedules, nil
}

func (r *postgresRepository) CreateSchedules(schedules []Schedule) error {
	if len(schedules) == 0 {
		return nil
	}

	// トランザクション開始
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `INSERT INTO schedules (id, user_id, game_id, start_time, end_time, status)
			  VALUES (?, ?, ?, ?, ?, ?)`

	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, schedule := range schedules {
		_, err := stmt.Exec(schedule.ID, schedule.UserID, schedule.GameID, schedule.StartTime, schedule.EndTime, schedule.Status)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *postgresRepository) UpdateScheduleStatus(scheduleID string, status string) error {
	query := `UPDATE schedules SET status = ? WHERE id = ?`
	_, err := r.db.Exec(query, status, scheduleID)
	return err
}

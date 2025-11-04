package calendar

import (
	"TO-DO-IT/internal/db" // (main.goで定義するDB)
)

// Repository ... データベース（今回はインメモリDB）とのインターフェース
type Repository interface {
	GetFixedEventsByUserID(userID string) ([]FixedEvent, error)
	CreateFixedEvent(event *FixedEvent) error
	GetSchedulesByUserID(userID string) ([]Schedule, error)
	CreateSchedules(schedules []Schedule) error
	UpdateScheduleStatus(scheduleID string, status ScheduleStatus) (*Schedule, error)
	DeleteSchedulesByUserID(userID string) error
}

type inMemoryRepository struct {
	db *db.MemoryDB
}

func NewRepository(db *db.MemoryDB) Repository {
	return &inMemoryRepository{db: db}
}

// --- FixedEvent (固定予定) ---
func (r *inMemoryRepository) GetFixedEventsByUserID(userID string) ([]FixedEvent, error) {
	r.db.RWMutex.RLock()
	defer r.db.RWMutex.RUnlock()

	var events []FixedEvent
	for _, event := range r.db.FixedEvents {
		if event.UserID == userID {
			events = append(events, event)
		}
	}
	return events, nil
}

func (r *inMemoryRepository) CreateFixedEvent(event *FixedEvent) error {
	r.db.RWMutex.Lock()
	defer r.db.RWMutex.Unlock()

	if _, exists := r.db.FixedEvents[event.ID]; exists {
		return db.ErrAlreadyExists
	}
	r.db.FixedEvents[event.ID] = *event
	return nil
}

// --- Schedule (自動生成スケジュール) ---
func (r *inMemoryRepository) GetSchedulesByUserID(userID string) ([]Schedule, error) {
	r.db.RWMutex.RLock()
	defer r.db.RWMutex.RUnlock()

	var schedules []Schedule
	for _, s := range r.db.Schedules {
		if s.UserID == userID {
			schedules = append(schedules, s)
		}
	}
	// (ソート処理を追加するのが望ましい)
	return schedules, nil
}

// CreateSchedules ... 複数のスケジュールを一括登録
func (r *inMemoryRepository) CreateSchedules(schedules []Schedule) error {
	r.db.RWMutex.Lock()
	defer r.db.RWMutex.Unlock()

	for _, s := range schedules {
		r.db.Schedules[s.ID] = s
	}
	return nil
}

// UpdateScheduleStatus ... スケジュールの進捗を更新
func (r *inMemoryRepository) UpdateScheduleStatus(scheduleID string, status ScheduleStatus) (*Schedule, error) {
	r.db.RWMutex.Lock()
	defer r.db.RWMutex.Unlock()

	s, exists := r.db.Schedules[scheduleID]
	if !exists {
		return nil, db.ErrNotFound
	}
	s.Status = status
	r.db.Schedules[scheduleID] = s
	return &s, nil
}

// DeleteSchedulesByUserID ... ユーザーの既存スケジュールを全削除 (自動生成の前に呼ぶ)
func (r *inMemoryRepository) DeleteSchedulesByUserID(userID string) error {
	r.db.RWMutex.Lock()
	defer r.db.RWMutex.Unlock()

	for id, s := range r.db.Schedules {
		if s.UserID == userID {
			delete(r.db.Schedules, id)
		}
	}
	return nil
}

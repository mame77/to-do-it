package calendar

import (
	"time"

	"github.com/mame77/to-do-it/backend/internal/db"
)

// Repository (インターフェース)
type Repository interface {
	// 固定予定 (FixedEvent)
	GetFixedEventsByUserID(userID string, start time.Time, end time.Time) ([]FixedEvent, error)
	CreateFixedEvent(event *FixedEvent) error

	// スケジュール (Schedule)
	GetSchedulesByUserID(userID string, start time.Time, end time.Time) ([]Schedule, error)
	CreateSchedules(schedules []Schedule) error
	UpdateScheduleStatus(scheduleID string, status string) error
}

// memoryRepository is an in-memory implementation using db.MemoryDB
type memoryRepository struct {
	db *db.MemoryDB
}

// NewRepository ... initialize repository with in-memory DB
func NewRepository(mdb *db.MemoryDB) Repository {
	return &memoryRepository{db: mdb}
}

func (r *memoryRepository) GetFixedEventsByUserID(userID string, start time.Time, end time.Time) ([]FixedEvent, error) {
	r.db.RWMutex.RLock()
	defer r.db.RWMutex.RUnlock()
	var events []FixedEvent
	for _, e := range r.db.FixedEvents {
		if e.UserID == userID && e.StartTime.Before(end) && e.EndTime.After(start) {
			events = append(events, e)
		}
	}
	return events, nil
}

func (r *memoryRepository) CreateFixedEvent(event *FixedEvent) error {
	r.db.RWMutex.Lock()
	defer r.db.RWMutex.Unlock()
	r.db.FixedEvents[event.ID] = *event
	return nil
}

func (r *memoryRepository) GetSchedulesByUserID(userID string, start time.Time, end time.Time) ([]Schedule, error) {
	r.db.RWMutex.RLock()
	defer r.db.RWMutex.RUnlock()
	var schedules []Schedule
	for _, s := range r.db.Schedules {
		if s.UserID == userID && s.StartTime.Before(end) && s.EndTime.After(start) {
			schedules = append(schedules, s)
		}
	}
	return schedules, nil
}

func (r *memoryRepository) CreateSchedules(schedules []Schedule) error {
	r.db.RWMutex.Lock()
	defer r.db.RWMutex.Unlock()
	for _, s := range schedules {
		r.db.Schedules[s.ID] = s
	}
	return nil
}

func (r *memoryRepository) UpdateScheduleStatus(scheduleID string, status string) error {
	r.db.RWMutex.Lock()
	defer r.db.RWMutex.Unlock()
	s, ok := r.db.Schedules[scheduleID]
	if !ok {
		return nil
	}
	s.Status = status
	r.db.Schedules[scheduleID] = s
	return nil
}

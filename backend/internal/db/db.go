package db

import (
	"errors"
	"sync"

	"github.com/mame77/to-do-it/backend/internal/model"
)

var (
	ErrNotFound      = errors.New("item not found")
	ErrAlreadyExists = errors.New("item already exists")
)

// MemoryDB ... アプリケーションのインメモリデータベース
type MemoryDB struct {
	RWMutex     sync.RWMutex
	Games       map[string]model.Game       // key: GameID
	FixedEvents map[string]model.FixedEvent // key: EventID
	Schedules   map[string]model.Schedule   // key: ScheduleID
	Motivations map[string]model.Motivation // key: UserID
}

func NewMemoryDB() *MemoryDB {
	return &MemoryDB{
		Games:       make(map[string]model.Game),
		FixedEvents: make(map[string]model.FixedEvent),
		Schedules:   make(map[string]model.Schedule),
		Motivations: make(map[string]model.Motivation),
	}
}

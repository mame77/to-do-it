package db

import (
	"TO-DO-IT/internal/calendar"
	"TO-DO-IT/internal/game"
	"TO-DO-IT/internal/score"
	"errors"
	"sync"
)

var (
	ErrNotFound      = errors.New("item not found")
	ErrAlreadyExists = errors.New("item already exists")
)

// MemoryDB ... アプリケーションのインメモリデータベース
type MemoryDB struct {
	RWMutex     sync.RWMutex
	Games       map[string]game.Game           // key: GameID
	FixedEvents map[string]calendar.FixedEvent // key: EventID
	Schedules   map[string]calendar.Schedule   // key: ScheduleID
	Motivations map[string]score.Motivation    // key: UserID
}

func NewMemoryDB() *MemoryDB {
	return &MemoryDB{
		Games:       make(map[string]game.Game),
		FixedEvents: make(map[string]calendar.FixedEvent),
		Schedules:   make(map[string]calendar.Schedule),
		Motivations: make(map[string]score.Motivation),
	}
}

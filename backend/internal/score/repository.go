package score

import "github.com/mame77/to-do-it/backend/internal/db"

type Repository interface {
	GetMotivationByUserID(userID string) (*Motivation, error)
	UpdateMotivation(motivation *Motivation) error
}

type memoryRepository struct {
	db *db.MemoryDB
}

func NewRepository(mdb *db.MemoryDB) Repository {
	return &memoryRepository{db: mdb}
}

func (r *memoryRepository) GetMotivationByUserID(userID string) (*Motivation, error) {
	r.db.RWMutex.RLock()
	defer r.db.RWMutex.RUnlock()
	m, ok := r.db.Motivations[userID]
	if !ok {
		return nil, nil
	}
	return &m, nil
}

func (r *memoryRepository) UpdateMotivation(motivation *Motivation) error {
	r.db.RWMutex.Lock()
	defer r.db.RWMutex.Unlock()
	r.db.Motivations[motivation.UserID] = *motivation
	return nil
}

package score

import (
	"TO-DO-IT/internal/db" // (main.goで定義するDB)
)

// Repository ... データベース（今回はインメモリDB）とのインターフェース
type Repository interface {
	GetMotivationByUserID(userID string) (Motivation, error)
	UpsertMotivation(motivation Motivation) error
}

type inMemoryRepository struct {
	db *db.MemoryDB
}

func NewRepository(db *db.MemoryDB) Repository {
	return &inMemoryRepository{db: db}
}

// GetMotivationByUserID ... モチベーション情報を取得
func (r *inMemoryRepository) GetMotivationByUserID(userID string) (Motivation, error) {
	r.db.RWMutex.RLock()
	defer r.db.RWMutex.RUnlock()

	m, exists := r.db.Motivations[userID]
	if !exists {
		// 存在しない場合はデフォルト値を返す
		return Motivation{UserID: userID, Points: 0, Rank: "ブロンズ", Level: 0}, nil
	}
	return m, nil
}

// UpsertMotivation ... モチベーション情報を更新 (なければ作成)
func (r *inMemoryRepository) UpsertMotivation(motivation Motivation) error {
	r.db.RWMutex.Lock()
	defer r.db.RWMutex.Unlock()

	r.db.Motivations[motivation.UserID] = motivation
	return nil
}

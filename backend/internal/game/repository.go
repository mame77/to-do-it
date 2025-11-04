package game

import "TO-DO-IT/internal/db" // (main.goで定義するDB)

// Repository ... データベース（今回はインメモリDB）とのインターフェース
type Repository interface {
	GetGamesByUserIDAndStatus(userID string, status Status) ([]Game, error)
	CreateGame(game *Game) error
}

type inMemoryRepository struct {
	db *db.MemoryDB
}

func NewRepository(db *db.MemoryDB) Repository {
	return &inMemoryRepository{db: db}
}

// GetGamesByUserIDAndStatus ... 指定ステータスのゲームを取得
func (r *inMemoryRepository) GetGamesByUserIDAndStatus(userID string, status Status) ([]Game, error) {
	r.db.RWMutex.RLock()
	defer r.db.RWMutex.RUnlock()

	var games []Game
	for _, game := range r.db.Games {
		if game.UserID == userID && game.Status == status {
			games = append(games, game)
		}
	}
	return games, nil
}

// CreateGame ... ゲームを登録
func (r *inMemoryRepository) CreateGame(game *Game) error {
	r.db.RWMutex.Lock()
	defer r.db.RWMutex.Unlock()

	if _, exists := r.db.Games[game.ID]; exists {
		return db.ErrAlreadyExists
	}
	r.db.Games[game.ID] = *game
	return nil
}

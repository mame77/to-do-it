package game

import (
	"fmt"
	"time"

	"github.com/mame77/to-do-it/backend/internal/db"
)

// Repository は、game データの永続化（DB操作）に関するインターフェースです。
type Repository interface {
	CreateGame(game *Game) (string, error)
	GetGameByID(id string) (*Game, error)
	GetGamesByUserID(userID string) ([]*Game, error)
	UpdateGame(game *Game) error
	DeleteGame(id string) error
}

// repository は Repository インターフェースの具体的な実装です。
// DB接続（*sql.DB）を持ちます。
type repository struct {
	db *db.MemoryDB
}

// NewRepository は、新しい repository インスタンスを作成します。
// main.go などでDB接続を確立した後、それを渡して呼び出します。
func NewRepository(mdb *db.MemoryDB) Repository {
	return &repository{db: mdb}
}

// --- インターフェースの実装 ---

// CreateGame は新しいゲームをDBに作成します。作成したゲームのIDを返します。
func (r *repository) CreateGame(game *Game) (string, error) {
	// in-memory 実装: ユニークIDを生成して map に格納
	id := fmt.Sprintf("game_%d", time.Now().UnixNano())
	now := time.Now()
	game.ID = id
	game.CreatedAt = now
	game.UpdatedAt = now

	r.db.RWMutex.Lock()
	r.db.Games[id] = *game
	r.db.RWMutex.Unlock()

	return id, nil
}

// GetGameByID は ID でゲームを1件取得します。
func (r *repository) GetGameByID(id string) (*Game, error) {
	r.db.RWMutex.RLock()
	g, ok := r.db.Games[id]
	r.db.RWMutex.RUnlock()
	if !ok {
		return nil, nil
	}
	return &g, nil
}

// GetGamesByUserID は、指定されたユーザーのゲーム一覧を取得します。
func (r *repository) GetGamesByUserID(userID string) ([]*Game, error) {
	r.db.RWMutex.RLock()
	defer r.db.RWMutex.RUnlock()

	var games []*Game
	for _, g := range r.db.Games {
		if g.UserID == userID {
			gCopy := g
			games = append(games, &gCopy)
		}
	}
	return games, nil
}

// UpdateGame はゲーム情報を更新します。
func (r *repository) UpdateGame(game *Game) error {
	r.db.RWMutex.Lock()
	defer r.db.RWMutex.Unlock()
	if _, ok := r.db.Games[game.ID]; !ok {
		return fmt.Errorf("game not found")
	}
	game.UpdatedAt = time.Now()
	r.db.Games[game.ID] = *game
	return nil
}

// DeleteGame は ID を指定してゲームを削除します。
func (r *repository) DeleteGame(id string) error {
	r.db.RWMutex.Lock()
	defer r.db.RWMutex.Unlock()
	delete(r.db.Games, id)
	return nil
}

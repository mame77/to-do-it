package game

import (
	"database/sql"
	"log"
	"time"
)

// Repository は、game データの永続化（DB操作）に関するインターフェースです。
type Repository interface {
	CreateGame(game *Game) (int, error)
	GetGameByID(id int) (*Game, error)
	GetGamesByUserID(userID int) ([]*Game, error)
	UpdateGame(game *Game) error
	DeleteGame(id int) error
}

// repository は Repository インターフェースの具体的な実装です。
// DB接続（*sql.DB）を持ちます。
type repository struct {
	db *sql.DB
}

// NewRepository は、新しい repository インスタンスを作成します。
// main.go などでDB接続を確立した後、それを渡して呼び出します。
func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

// --- インターフェースの実装 ---

// CreateGame は新しいゲームをDBに作成します。作成したゲームのIDを返します。
func (r *repository) CreateGame(game *Game) (int, error) {
	// 認証なしの暫定対応として、game.UserID はサービス層で設定済みと仮定
	query := `INSERT INTO games (user_id, title, platform, genre, status, release_date, created_at, updated_at)
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	// Go 1.22以降なら time.Now() でOK。それ以前なら time.Now().UTC() などDBの型に合わせる
	now := time.Now()

	result, err := r.db.Exec(query,
		game.UserID,
		game.Title,
		game.Platform,
		game.Genre,
		game.Status,
		game.ReleaseDate,
		now, // CreatedAt
		now, // UpdatedAt
	)
	if err != nil {
		log.Printf("Error creating game: %v", err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert ID: %v", err)
		return 0, err
	}

	return int(id), nil
}

// GetGameByID は ID でゲームを1件取得します。
func (r *repository) GetGameByID(id int) (*Game, error) {
	query := `SELECT id, user_id, title, platform, genre, status, release_date, created_at, updated_at
			  FROM games WHERE id = ?`

	var game Game
	err := r.db.QueryRow(query, id).Scan(
		&game.ID,
		&game.UserID,
		&game.Title,
		&game.Platform,
		&game.Genre,
		&game.Status,
		&game.ReleaseDate,
		&game.CreatedAt,
		&game.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // 見つからなかった（エラーではない）
		}
		log.Printf("Error scanning game by ID: %v", err)
		return nil, err
	}

	return &game, nil
}

// GetGamesByUserID は、指定されたユーザーのゲーム一覧を取得します。
func (r *repository) GetGamesByUserID(userID int) ([]*Game, error) {
	query := `SELECT id, user_id, title, platform, genre, status, release_date, created_at, updated_at
			  FROM games WHERE user_id = ? ORDER BY created_at DESC`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		log.Printf("Error querying games by user ID: %v", err)
		return nil, err
	}
	defer rows.Close()

	var games []*Game
	for rows.Next() {
		var game Game
		if err := rows.Scan(
			&game.ID,
			&game.UserID,
			&game.Title,
			&game.Platform,
			&game.Genre,
			&game.Status,
			&game.ReleaseDate,
			&game.CreatedAt,
			&game.UpdatedAt,
		); err != nil {
			log.Printf("Error scanning game row: %v", err)
			continue // 一部の行でエラーがあっても続行
		}
		games = append(games, &game)
	}

	return games, nil
}

// UpdateGame はゲーム情報を更新します。
func (r *repository) UpdateGame(game *Game) error {
	query := `UPDATE games SET title = ?, platform = ?, genre = ?, status = ?, release_date = ?, updated_at = ?
			  WHERE id = ?`

	_, err := r.db.Exec(query,
		game.Title,
		game.Platform,
		game.Genre,
		game.Status,
		game.ReleaseDate,
		time.Now(), // UpdatedAt
		game.ID,
	)
	
	if err != nil {
		log.Printf("Error updating game: %v", err)
	}
	return err
}

// DeleteGame は ID を指定してゲームを削除します。
func (r *repository) DeleteGame(id int) error {
	query := `DELETE FROM games WHERE id = ?`
	
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting game: %v", err)
	}
	return err
}
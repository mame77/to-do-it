package game

import (
	"log"
)

// testUserID は認証をスキップするための仮のユーザーID
const testUserID = 1 

// Service は、game のビジネスロジックに関するインターフェースです。
type Service interface {
	// 認証がないため、UserIDはサービス内で固定値(1)を使います
	CreateGame(req *CreateGameRequest) (*Game, error)
	GetGame(id int) (*Game, error)
	GetGames() ([]*Game, error) // UserIDを引数に取らず、固定値(1)で検索
	UpdateGame(id int, req *UpdateGameRequest) (*Game, error)
	DeleteGame(id int) error
}

// service は Service インターフェースの具体的な実装です。
// repository（DB操作）を持ちます。
type service struct {
	repo Repository
}

// NewService は、新しい service インスタンスを作成します。
// handler が repository を渡して呼び出します。
func NewService(repo Repository) Service {
	return &service{repo: repo}
}

// --- インターフェースの実装 ---

// CreateGame は新しいゲームを作成します。
func (s *service) CreateGame(req *CreateGameRequest) (*Game, error) {
	// リクエスト(Request)からDBモデル(Game)へ変換
	status := req.Status
	if status == "" {
		status = "unstarted" // デフォルト値
	}
	game := &Game{
		UserID:      testUserID, // ★認証の代わりに固定IDを設定
		Title:       req.Title,
		Platform:    req.Platform,
		Genre:       req.Genre,
		Status:      status,
		ReleaseDate: req.ReleaseDate,
		// CreatedAt/UpdatedAt は repository 層のSQLで設定
	}

	// リポジトリを呼び出してDBに保存
	id, err := s.repo.CreateGame(game)
	if err != nil {
		log.Printf("Service: Error creating game: %v", err)
		return nil, err
	}

	// DBに保存された最新の情報を取得して返す（IDなどが確定するため）
	// ※CreateGame が Game オブジェクトを丸ごと返せば不要だが、今回はIDのみ返すと仮定
	createdGame, err := s.repo.GetGameByID(id)
	if err != nil {
		log.Printf("Service: Error fetching created game: %v", err)
		return nil, err
	}
	
	return createdGame, nil
}

// GetGame は ID でゲームを1件取得します。
func (s *service) GetGame(id int) (*Game, error) {
	game, err := s.repo.GetGameByID(id)
	if err != nil {
		log.Printf("Service: Error getting game by ID: %v", err)
		return nil, err
	}
	if game == nil {
		// 見つからない
		return nil, nil // エラーではなく nil を返す
	}

	// TODO: 本来はここで「取得した game.UserID」と「認証ユーザーID」が一致するかチェックする
	// if game.UserID != testUserID {
	// 	return nil, errors.New("forbidden") // 他人のゲーム
	// }

	return game, nil
}

// GetGames は（テストユーザーの）ゲーム一覧を取得します。
func (s *service) GetGames() ([]*Game, error) {
	// 認証の代わりに固定IDで検索
	games, err := s.repo.GetGamesByUserID(testUserID)
	if err != nil {
		log.Printf("Service: Error getting games by UserID: %v", err)
		return nil, err
	}
	return games, nil
}

// UpdateGame はゲーム情報を更新します。
func (s *service) UpdateGame(id int, req *UpdateGameRequest) (*Game, error) {
	// 1. まず対象のゲームが存在するか確認
	game, err := s.repo.GetGameByID(id)
	if err != nil {
		return nil, err
	}
	if game == nil {
		return nil, nil // 見つからない
	}

	// TODO: 本来はここで「取得した game.UserID」と「認証ユーザーID」が一致するかチェックする
	// if game.UserID != testUserID {
	// 	return nil, errors.New("forbidden") // 他人のゲーム
	// }

	// 2. リクエスト(req)の内容で、取得した game オブジェクトを更新
	// ※リクエストで値が省略された場合（例：Title=""）にどうするかは要件次第
	// ここでは単純に上書きする
	if req.Title != "" {
		game.Title = req.Title
	}
	if req.Platform != "" {
		game.Platform = req.Platform
	}
	if req.Genre != "" {
		game.Genre = req.Genre
	}
	if req.Status != "" {
		game.Status = req.Status
	}
	if !req.ReleaseDate.IsZero() {
		game.ReleaseDate = req.ReleaseDate
	}
	// UpdatedAt は repository 層で更新

	// 3. DBを更新
	err = s.repo.UpdateGame(game)
	if err != nil {
		log.Printf("Service: Error updating game: %v", err)
		return nil, err
	}
	
	return game, nil
}

// DeleteGame は ID を指定してゲームを削除します。
func (s *service) DeleteGame(id int) error {
	// 1. まず対象のゲームが存在するか確認（しなくてもDBエラーにはなるが、権限チェックのため）
	game, err := s.repo.GetGameByID(id)
	if err != nil {
		return err
	}
	if game == nil {
		return nil // 削除対象が見つからない（エラーではない）
	}
	
	// TODO: 本来はここで「取得した game.UserID」と「認証ユーザーID」が一致するかチェックする
	// if game.UserID != testUserID {
	// 	return errors.New("forbidden") // 他人のゲーム
	// }

	// 2. 削除実行
	return s.repo.DeleteGame(id)
}
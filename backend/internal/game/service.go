package game

import "TO-DO-IT/internal/app" // (main.goで定義するヘルパー)

// Service ... ビジネスロジック
type Service interface {
	CreateGame(userID, title, genre string) (*Game, error)
	GetPendingGames(userID string) ([]Game, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateGame(userID, title, genre string) (*Game, error) {
	game := &Game{
		ID:     app.NewUUID(),
		UserID: userID,
		Title:  title,
		Genre:  genre,
		Status: StatusPending,
	}
	if err := s.repo.CreateGame(game); err != nil {
		return nil, err
	}
	return game, nil
}

func (s *service) GetPendingGames(userID string) ([]Game, error) {
	return s.repo.GetGamesByUserIDAndStatus(userID, StatusPending)
}

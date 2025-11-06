package score

// Service (インターフェース)
type Service interface {
	GetMotivation(userID string) (*Motivation, error)
	ReportPlayResult(userID string, result PlayResult) (*Motivation, error) // [cite: 76]
}

type service struct {
	repo Repository
	// taskRepo task.Repository // (もしプレイセッション[cite: 96-101]の保存が必要なら)
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetMotivation(userID string) (*Motivation, error) {
	return s.repo.GetMotivationByUserID(userID)
}

// ReportPlayResult (ボーナス・ペナルティロジック)
func (s *service) ReportPlayResult(userID string, result PlayResult) (*Motivation, error) {
	// 1. 現在のモチベーションを取得
	motivation, err := s.repo.GetMotivationByUserID(userID)
	if err != nil {
		return nil, err
	}

	// 2. 結果に応じてポイントを増減 (仮のロジック) [cite: 77-78]
	if result.Result == "success" {
		motivation.Points += 10 // ボーナス
	} else {
		motivation.Points -= 5 // ペナルティ
	}

	// TODO: ポイントに応じてランクやレベルを更新するロジック

	// 3. DBに保存
	if err := s.repo.UpdateMotivation(motivation); err != nil {
		return nil, err
	}

	// (TODO: taskRepoを使ってplay_sessions [cite: 96-101] にプレイ記録を保存する)

	return motivation, nil
}

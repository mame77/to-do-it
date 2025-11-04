package score

// Service ... ビジネスロジック
type Service interface {
	GetMotivation(userID string) (Motivation, error)
	ReportPlayResult(userID string, result PlayResult) (Motivation, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetMotivation(userID string) (Motivation, error) {
	return s.repo.GetMotivationByUserID(userID)
}

[cite_start]// ReportPlayResult ... プレイ結果を反映し、ボーナス・ペナルティを計算 [cite: 150-154]
func (s *service) ReportPlayResult(userID string, result PlayResult) (Motivation, error) {
	// 1. 現在のモチベーションを取得
	m, err := s.repo.GetMotivationByUserID(userID)
	if err != nil {
		return m, err
	}

	// 2. 結果に応じてポイントを増減 (ロジック)
	if result.Result == "success" {
		[cite_start]m.Points += 10 // ボーナス [cite: 26]
	} else {
		[cite_start]m.Points -= 5 // ペナルティ [cite: 25]
		if m.Points < 0 {
			m.Points = 0
		}
	}

	// 3. ランクを更新 (ロジック)
	if m.Points > 100 {
		m.Rank = "シルバー"
	} else {
		[cite_start]m.Rank = "ブロンズ" // [cite: 43]
	}
	m.Level = m.Points % 100 // (レベル = ポイントの100の剰余 とする)

	// 4. DBに保存
	if err := s.repo.UpsertMotivation(m); err != nil {
		return m, err
	}
	return m, nil
}
package calendar

// Repository はデータ永続化層へのインターフェースを定義します。
// 実際にはデータベースや外部ストレージへの接続コードが含まれます。
type Repository interface {
	// Game
	GetGamesToSchedule() ([]Game, error)
	GetGameByID(id string) (*Game, error)

	// Fixed Event
	GetAllFixedEvents() ([]FixedEvent, error)

	// Schedule
	SaveSchedules(schedules []Schedule) error
	UpdateScheduleStatus(id string, completed bool, skipped bool) error
}

// NewRepository は実際のRepository実装を初期化して返します。
// (ここではダミーの実装を返す)
func NewRepository() Repository {
	return &mockRepository{} // 開発用モックまたは実際のDB実装
}

// --- ダミーのRepository実装 (開発・テスト用) ---
type mockRepository struct{}

func (r *mockRepository) GetGamesToSchedule() ([]Game, error) {
	// ダミーデータ: 未完了かプレイ中のゲーム
	return []Game{
		{ID: "g1", Title: "ゼルダの伝説", Status: StatusPlaying, AddedAt: "2024-01-01T00:00:00Z"},
		{ID: "g2", Title: "ファイナルファンタジーXIV", Status: StatusPlaying, AddedAt: "2024-02-01T00:00:00Z"},
	}, nil
}

func (r *mockRepository) GetGameByID(id string) (*Game, error) {
	// IDに基づいたゲーム取得ロジック
	return &Game{ID: id, Title: "テストゲーム", Status: StatusPlaying, AddedAt: "2024-01-01T00:00:00Z"}, nil
}

func (r *mockRepository) GetAllFixedEvents() ([]FixedEvent, error) {
	// ダミーデータ: 毎週月水金の仕事
	return []FixedEvent{
		{
			ID: "fe1",
			Title: "仕事",
			DayOfWeek: []int{1, 3, 5}, // 月水金
			StartTime: "09:00",
			EndTime: "18:00",
			IsRecurring: true,
		},
	}, nil
}

func (r *mockRepository) SaveSchedules(schedules []Schedule) error {
	// スケジュールを保存するダミーロジック
	return nil
}

func (r *mockRepository) UpdateScheduleStatus(id string, completed bool, skipped bool) error {
	// スケジュールステータスを更新するダミーロジック
	return nil
}
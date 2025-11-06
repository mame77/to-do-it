package calendar

import (
	"time"

	"github.com/mame77/to-do-it/backend/internal/game" // 担当Cのゲームパッケージ (仮)
	"github.com/mame77/to-do-it/backend/internal/score"
)

// Service (インターフェース)
type Service interface {
	// 自動生成ロジック [cite: 71]
	GenerateSchedule(userID string) ([]Schedule, error)

	// スケジュール取得
	GetSchedules(userID string, start time.Time, end time.Time) ([]Schedule, error)
	// スケジュール進捗更新
	UpdateScheduleStatus(scheduleID string, status string) error

	// 固定予定
	GetFixedEvents(userID string, start time.Time, end time.Time) ([]FixedEvent, error)
	CreateFixedEvent(event *FixedEvent) error
}

// service (実装)
type service struct {
	calendarRepo Repository
	gameRepo     game.Repository // 担当Cのゲームリポジトリ (仮)
	scoreSvc     score.Service
}

// NewService ... 必要なリポジトリを受け取り、サービスを初期化
func NewService(calRepo Repository, gameRepo game.Repository, scoreSvc score.Service) Service {
	return &service{
		calendarRepo: calRepo,
		gameRepo:     gameRepo,
		scoreSvc:     scoreSvc,
	}
}

// GenerateSchedule (最重要ロジック)
func (s *service) GenerateSchedule(userID string) ([]Schedule, error) {
	// --- ここに自動生成アルゴリズムを実装します [cite: 71, 124] ---

	// 1. 担当Cのリポジトリから、ユーザーの「未開始」ゲームを取得 [cite: 40]
	// games, err := s.gameRepo.GetGamesByUserID(userID, "未開始")
	// if err != nil { ... }

	// 2. このリポジトリから、ユーザーの「固定予定」を取得 [cite: 126, 131-132]
	// fixedEvents, err := s.calendarRepo.GetFixedEventsByUserID(userID, time.Now(), time.Now().Add(7*24*time.Hour)) // 例えば1週間分
	// if err != nil { ... }

	// 3. 空き時間を計算する
	// (固定予定を考慮して、ゲームをプレイできる時間帯を算出するロジック)

	// 4. 空き時間にゲームを割り当てる (仮のロジック)
	var newSchedules []Schedule
	// ... (アルゴリズム) ...
	// newSchedules = append(newSchedules, Schedule{...})

	// 5. 生成したスケジュールをDBに保存
	// if err := s.calendarRepo.CreateSchedules(newSchedules); err != nil { ... }

	return newSchedules, nil
}

func (s *service) GetSchedules(userID string, start time.Time, end time.Time) ([]Schedule, error) {
	return s.calendarRepo.GetSchedulesByUserID(userID, start, end)
}

func (s *service) UpdateScheduleStatus(scheduleID string, status string) error {
	// TODO: ステータス更新時に、scoreパッケージ(担当A)のサービスを呼び出し、
	// ボーナス・ペナルティを発生させる必要がある [cite: 76]
	// if status == "完了" { s.scoreService.ReportPlayResult(scheduleID, "success") }
	return s.calendarRepo.UpdateScheduleStatus(scheduleID, status)
}

func (s *service) GetFixedEvents(userID string, start time.Time, end time.Time) ([]FixedEvent, error) {
	return s.calendarRepo.GetFixedEventsByUserID(userID, start, end)
}

func (s *service) CreateFixedEvent(event *FixedEvent) error {
	// TODO: バリデーションチェック (時間が重複していないか等)
	return s.calendarRepo.CreateFixedEvent(event)
}

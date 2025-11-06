package calendar

import (
	"TO-DO-IT/internal/game" // 担当Cのゲームパッケージ (仮)
	"fmt"
	"strconv"
	"time"
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
}

// NewService ... 必要なリポジトリを受け取り、サービスを初期化
func NewService(calRepo Repository, gameRepo game.Repository) Service {
	return &service{
		calendarRepo: calRepo,
		gameRepo:     gameRepo,
	}
}

// GenerateSchedule (最重要ロジック)
func (s *service) GenerateSchedule(userID string) ([]Schedule, error) {
	// 1. ユーザーの「未開始」ゲームを取得 (固定UserID=1を使用)
	games, err := s.gameRepo.GetGamesByUserID(1) // testUserID = 1
	if err != nil {
		return nil, err
	}

	// 未開始のゲームのみをフィルタリング
	var unstartedGames []*game.Game
	for _, g := range games {
		if g.Status == "unstarted" {
			unstartedGames = append(unstartedGames, g)
		}
	}

	if len(unstartedGames) == 0 {
		return []Schedule{}, nil // 未開始ゲームがない場合は空を返す
	}

	// 2. 今後1週間分の固定予定を取得
	start := time.Now()
	end := start.Add(7 * 24 * time.Hour)
	fixedEvents, err := s.calendarRepo.GetFixedEventsByUserID(userID, start, end)
	if err != nil {
		return nil, err
	}

	// 3. スケジュールを生成（シンプルなアルゴリズム）
	var newSchedules []Schedule
	currentTime := start

	for idx, g := range unstartedGames {
		// 各ゲームに2時間のプレイ時間を割り当て
		playDuration := 2 * time.Hour

		// 固定予定と重ならない時間を探す
		scheduleTime := findNextAvailableTime(currentTime, playDuration, fixedEvents)

		// スケジュールを作成
		schedule := Schedule{
			ID:        generateScheduleID(idx),
			UserID:    userID,
			GameID:    strconv.Itoa(g.ID), // intをstringに変換
			StartTime: scheduleTime,
			EndTime:   scheduleTime.Add(playDuration),
			Status:    "pending",
		}
		newSchedules = append(newSchedules, schedule)

		// 次のスケジュールは現在のスケジュール終了後から
		currentTime = schedule.EndTime
	}

	// 4. 生成したスケジュールをDBに保存
	if len(newSchedules) > 0 {
		if err := s.calendarRepo.CreateSchedules(newSchedules); err != nil {
			return nil, err
		}
	}

	return newSchedules, nil
}

// findNextAvailableTime は固定予定と重ならない次の利用可能な時間を見つける
func findNextAvailableTime(startTime time.Time, duration time.Duration, fixedEvents []FixedEvent) time.Time {
	proposedTime := startTime

	// 営業時間内に調整（9:00-23:00）
	proposedTime = adjustToBusinessHours(proposedTime)

	for {
		proposedEnd := proposedTime.Add(duration)
		conflict := false

		// 固定予定との衝突をチェック
		for _, event := range fixedEvents {
			if timeOverlaps(proposedTime, proposedEnd, event.StartTime, event.EndTime) {
				conflict = true
				// 固定予定の終了後に移動
				proposedTime = event.EndTime
				proposedTime = adjustToBusinessHours(proposedTime)
				break
			}
		}

		if !conflict {
			return proposedTime
		}
	}
}

// timeOverlaps は2つの時間範囲が重なっているかチェック
func timeOverlaps(start1, end1, start2, end2 time.Time) bool {
	return start1.Before(end2) && end1.After(start2)
}

// adjustToBusinessHours は時間を営業時間内（9:00-23:00）に調整
func adjustToBusinessHours(t time.Time) time.Time {
	hour := t.Hour()

	// 23時以降なら翌日9時に
	if hour >= 23 || hour < 9 {
		t = time.Date(t.Year(), t.Month(), t.Day()+1, 9, 0, 0, 0, t.Location())
	}

	return t
}

// generateScheduleID はユニークなスケジュールIDを生成
func generateScheduleID(index int) string {
	return fmt.Sprintf("sched_%s_%d_%d", time.Now().Format("20060102150405"), time.Now().Nanosecond(), index)
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

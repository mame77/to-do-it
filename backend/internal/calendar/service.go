package calendar

import (
	"fmt"
	"time"
)

// Service はカレンダー機能のビジネスロジックを定義します。
type Service struct {
	repo Repository
}

// NewService はServiceを初期化します。
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// GenerateSchedules はゲームと固定予定に基づいてスケジュールを自動生成します。
func (s *Service) GenerateSchedules(req GenerateScheduleRequest) ([]Schedule, error) {
	games, err := s.repo.GetGamesToSchedule()
	if err != nil {
		return nil, fmt.Errorf("failed to get games: %w", err)
	}

	fixedEvents, err := s.repo.GetAllFixedEvents()
	if err != nil {
		return nil, fmt.Errorf("failed to get fixed events: %w", err)
	}

	// フロントエンドのロジック (lib/scheduleGenerator.ts) に相当する部分を呼び出し
	generated := generateScheduleCore(games, fixedEvents, req.StartDate, req.DaysToSchedule)
	
	// 生成されたスケジュールを永続化（例：DBに保存）
	if err := s.repo.SaveSchedules(generated); err != nil {
		return nil, fmt.Errorf("failed to save schedules: %w", err)
	}

	return generated, nil
}

// UpdateScheduleAction はスケジュールの完了またはスキップを処理します。
func (s *Service) UpdateScheduleAction(req ScheduleActionRequest) error {
	var completed bool
	var skipped bool

	switch req.Action {
	case "complete":
		completed = true
	case "skip":
		skipped = true
	default:
		return fmt.Errorf("invalid action: %s", req.Action)
	}

	// データベースでステータスを更新
	return s.repo.UpdateScheduleStatus(req.ScheduleID, completed, skipped)
}

// --- generateScheduleCore: スケジュール生成の主要ロジック（lib/scheduleGenerator.tsを再現） ---
func generateScheduleCore(games []Game, events []FixedEvent, startDateStr string, days int) []Schedule {
	schedules := []Schedule{}
	if len(games) == 0 {
		return schedules
	}

	sessionDuration := 90 // 90分セッション
	gameIndex := 0

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		startDate = time.Now()
	}

	for dayOffset := 0; dayOffset < days; dayOffset++ {
		currentDate := startDate.AddDate(0, 0, dayOffset)
		dateStr := currentDate.Format("2006-01-02")
		
		// 平日(1)か週末(3)かで目標セッション数を決定
		dayOfWeek := currentDate.Weekday()
		targetSessions := 1
		if dayOfWeek == time.Saturday || dayOfWeek == time.Sunday {
			targetSessions = 3
		}

		// findAvailableSlots のロジックを内部で実行（省略）
		// ここでは簡略化のため、ダミーのスロットを生成
		availableSlots := getDummySlots(currentDate)

		sessionsScheduled := 0
		for _, slot := range availableSlots {
			if sessionsScheduled >= targetSessions {
				break
			}
			
			// スロット内のプレイ可能な回数を計算
			slotDuration := timeToMinutes(slot.EndTime) - timeToMinutes(slot.StartTime)
			possibleSessions := slotDuration / sessionDuration
			
			currentStartMinutes := timeToMinutes(slot.StartTime)

			for i := 0; i < possibleSessions && sessionsScheduled < targetSessions; i++ {
				startMinutes := currentStartMinutes + (i * sessionDuration)
				endMinutes := startMinutes + sessionDuration

				game := games[gameIndex%len(games)]

				schedules = append(schedules, Schedule{
					ID:        fmt.Sprintf("schedule_%s_%d", dateStr, len(schedules)),
					GameID:    game.ID,
					Date:      dateStr,
					StartTime: minutesToTime(startMinutes),
					EndTime:   minutesToTime(endMinutes),
					Completed: false,
					Skipped:   false,
				})

				gameIndex++
				sessionsScheduled++
			}
		}
	}

	return schedules
}

// 以下はヘルパー関数 (lib/scheduleGenerator.ts の Go実装に相当)
func timeToMinutes(timeStr string) int {
	t, _ := time.Parse("15:04", timeStr)
	return t.Hour()*60 + t.Minute()
}

func minutesToTime(minutes int) string {
	hours := minutes / 60
	mins := minutes % 60
	return fmt.Sprintf("%02d:%02d", hours, mins)
}

// 開発用ダミー関数
func getDummySlots(date time.Time) []struct {StartTime string; EndTime string} {
	// 10:00-12:00, 14:00-16:00, 19:00-23:00 が空いていると仮定
	return []struct {StartTime string; EndTime string}{
		{"10:00", "12:00"},
		{"14:00", "16:00"},
		{"19:00", "23:00"},
	}
}
package calendar

import (
	"TO-DO-IT/internal/app"  // (main.goで定義するヘルパー)
	"TO-DO-IT/internal/game" // 担当Cのパッケージ
	"TO-DO-IT/internal/score" // 担当A(あなた)の別パッケージ
	"errors"
	"time"
)

// Service ... ビジネスロジック
type Service interface {
	[cite_start]GenerateSchedule(userID string) ([]Schedule, error) // [cite: 21]
	GetSchedules(userID string) ([]Schedule, error)
	[cite_start]UpdateScheduleStatus(scheduleID, status string) error // [cite: 22]
	[cite_start]CreateFixedEvent(userID, title string, start time.Time, end time.Time) (*FixedEvent, error) // [cite: 28]
	GetFixedEvents(userID string) ([]FixedEvent, error)
}

type service struct {
	calendarRepo Repository
	gameRepo     game.Repository // 担当Cのリポジトリ
	scoreSvc     score.Service   // 担当A(あなた)の別サービス
}

func NewService(calRepo Repository, gameRepo game.Repository, scoreSvc score.Service) Service {
	return &service{
		calendarRepo: calRepo,
		gameRepo:     gameRepo,
		scoreSvc:     scoreSvc,
	}
}

[cite_start]// GenerateSchedule ... スケジュール自動生成 [cite: 21]
func (s *service) GenerateSchedule(userID string) ([]Schedule, error) {
	[cite_start]// 1. 担当Cのサービスから「未開始」のゲームを取得 [cite: 19]
	pendingGames, err := s.gameRepo.GetGamesByUserIDAndStatus(userID, game.StatusPending)
	if err != nil {
		return nil, err
	}
	if len(pendingGames) == 0 {
		return nil, errors.New("スケジュールする未開始のゲームがありません")
	}

	// 2. 既存のスケジュールを全削除
	if err := s.calendarRepo.DeleteSchedulesByUserID(userID); err != nil {
		return nil, err
	}

	[cite_start]// 3. 固定予定を取得 [cite: 23, 28]
	fixedEvents, err := s.calendarRepo.GetFixedEventsByUserID(userID)
	if err != nil {
		return nil, err
	}

	// 4. スケジューリングロジック (簡易版)
	// 翌日から7日間、毎日午前9時に1時間のスロットを確保しようとする
	var newSchedules []Schedule
	gameIndex := 0
	now := time.Now()

	for i := 1; i <= 7; i++ { // 7日分
		if gameIndex >= len(pendingGames) {
			break // 全てのゲームを割り当てたら終了
		}

		// ターゲット時間: 翌日の午前9時、翌々日の午前9時...
		targetStart := time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, time.Local).AddDate(0, 0, i)
		targetEnd := targetStart.Add(1 * time.Hour)

		[cite_start]// 5. 固定予定と重複チェック [cite: 23]
		isOverlap := false
		for _, fe := range fixedEvents {
			// (targetStart < fe.EndTime) AND (targetEnd > fe.StartTime)
			if targetStart.Before(fe.EndTime) && targetEnd.After(fe.StartTime) {
				isOverlap = true
				break
			}
		}

		// 6. 重複がなければスケジュール作成
		if !isOverlap {
			game := pendingGames[gameIndex]
			schedule := Schedule{
				ID:        app.NewUUID(),
				UserID:    userID,
				GameID:    game.ID,
				GameTitle: game.Title,
				StartTime: targetStart,
				EndTime:   targetEnd,
				Status:    StatusScheduled,
			}
			newSchedules = append(newSchedules, schedule)
			gameIndex++
		}
	}

	// 7. DBに保存
	if err := s.calendarRepo.CreateSchedules(newSchedules); err != nil {
		return nil, err
	}

	return newSchedules, nil
}

// GetSchedules ... スケジュール一覧取得
func (s *service) GetSchedules(userID string) ([]Schedule, error) {
	return s.calendarRepo.GetSchedulesByUserID(userID)
}

// UpdateScheduleStatus ... スケジュールの進捗更新
func (s *service) UpdateScheduleStatus(scheduleID, statusStr string) error {
	var status ScheduleStatus
	var result string

	if statusStr == "completed" {
		status = StatusCompleted
		[cite_start]result = "success" // [cite: 152]
	} else if statusStr == "skipped" {
		status = StatusSkipped
		[cite_start]result = "failure" // [cite: 153]
	} else {
		return errors.New("無効なステータスです")
	}

	// 1. スケジュールDBを更新
	schedule, err := s.calendarRepo.UpdateScheduleStatus(scheduleID, status)
	if err != nil {
		return err
	}

	[cite_start]// 2. scoreサービスを呼び出し、ボーナス・ペナルティを反映 [cite: 25, 26, 150-151]
	playResult := score.PlayResult{
		ScheduleID: scheduleID,
		Result:     result,
	}
	// (ここでUserIDを渡す)
	if _, err := s.scoreSvc.ReportPlayResult(schedule.UserID, playResult); err != nil {
		return err // (エラーハンドリング: 本来は片方失敗してもOKにするか検討)
	}

	return nil
}

// --- 固定予定 ---

func (s *service) CreateFixedEvent(userID, title string, start time.Time, end time.Time) (*FixedEvent, error) {
	if start.After(end) {
		return nil, errors.New("開始時刻は終了時刻より前にしてください")
	}
	event := &FixedEvent{
		ID:        app.NewUUID(),
		UserID:    userID,
		Title:     title,
		StartTime: start,
		EndTime:   end,
	}
	if err := s.calendarRepo.CreateFixedEvent(event); err != nil {
		return nil, err
	}
	return event, nil
}

func (s *service) GetFixedEvents(userID string) ([]FixedEvent, error) {
	return s.calendarRepo.GetFixedEventsByUserID(userID)
}
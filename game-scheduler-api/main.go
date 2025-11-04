package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// --- 企画書から読み取れるデータモデル ---
// ※本来はDBから取得しますが、今回はモック（仮データ）として使います。

// Game (担当Cが管理するが、ロジックで必要)
type Game struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"` // 未開始, プレイ中, クリア済み
}

// --- 担当Aが定義・管理する必要があるデータ構造 (仮定義) ---

// FixedEvent (固定予定)
type FixedEvent struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"` // "仕事", "授業" など
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

// Schedule (自動生成されたスケジュール)
type Schedule struct {
	ID        string    `json:"id"`
	GameID    string    `json:"game_id"` // 紐づくゲームのID
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Status    string    `json:"status"` // "予定", "完了", "スキップ" など
}

// Motivation (継続状態)
type Motivation struct {
	Points int    `json:"points"`
	Rank   string `json:"rank"` // "ブロンズ" など
}

// PlayResult (プレイ結果リクエスト)
type PlayResult struct {
	ScheduleID string `json:"schedule_id"`
	Result     string `json:"result"` // "completed" (完了), "skipped" (失敗)
}

// --- メイン関数 (サーバー起動) ---

func main() {
	e := echo.New()

	// 担当AのAPIエンドポイントをグループ化
	api := e.Group("/api")
	{
		// 1. 手動スケジュール管理API (固定予定) [cite: 155-159]
		fixedApi := api.Group("/fixed-events")
		fixedApi.POST("", createFixedEvent)
		fixedApi.GET("", getFixedEvents)
		// PUT, DELETE もここに追加

		// 2. 自動スケジュール生成API [cite: 146]
		scheduleApi := api.Group("/schedule")
		scheduleApi.POST("/generate", generateSchedule) // [cite: 146]

		// 3. 生成済みスケジュールAPI [cite: 147-149]
		scheduleApi.GET("", getSchedules)        // [cite: 147]
		scheduleApi.PUT("/:id", updateSchedule) // [cite: 148]

		// 4. 継続促進ロジックAPI [cite: 150-154]
		motivationApi := api.Group("/motivation")
		motivationApi.GET("", getMotivationStatus)        // [cite: 154]
		motivationApi.POST("/result", reportPlayResult) // [cite: 151]
	}

	e.Logger.Fatal(e.Start(":8080"))
}

// --- 担当AのAPIハンドラ (スタブ) ---
// ここにあなたのロジックを実装していきます。

// 1. 固定予定の作成 [cite: 156]
func createFixedEvent(c echo.Context) error {
	fe := new(FixedEvent)
	if err := c.Bind(fe); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	// ★ TODO: ここでDBに保存する処理を実装 ★

	fe.ID = "fixed_event_123" // 仮のID
	return c.JSON(http.StatusCreated, fe)
}

// 1. 固定予定の一覧取得 [cite: 157]
func getFixedEvents(c echo.Context) error {
	// ★ TODO: ここでDBから固定予定一覧を取得する処理を実装 ★

	// 仮のデータを返す
	mockEvents := []FixedEvent{
		{ID: "fe_1", Title: "仕事", StartTime: time.Now(), EndTime: time.Now().Add(8 * time.Hour)},
		{ID: "fe_2", Title: "授業", StartTime: time.Now().Add(9 * time.Hour), EndTime: time.Now().Add(10 * time.Hour)},
	}
	return c.JSON(http.StatusOK, mockEvents)
}

// 2. スケジュール自動生成 [cite: 146]
func generateSchedule(c echo.Context) error {
	// ★ TODO: ここが最重要ロジック ★
	// 1. DBから「未開始」のゲーム一覧を取得 (担当CのAPI/DBをコール)
	// 2. DBから「固定予定」一覧を取得 (getFixedEvents のロジック)
	// 3. 1と2を元に、空いている時間にプレイ計画を生成するアルゴリズムを実装
	// 4. 生成したスケジュールをDBに保存

	return c.JSON(http.StatusOK, map[string]string{"message": "スケジュールを自動生成しました"})
}

// 3. 生成済みスケジュールの取得 [cite: 147]
func getSchedules(c echo.Context) error {
	// ★ TODO: DBから生成済みのスケジュール一覧を取得 ★

	// 仮のデータを返す
	mockSchedules := []Schedule{
		{ID: "sch_1", GameID: "game_abc", StartTime: time.Now().Add(24 * time.Hour), EndTime: time.Now().Add(25 * time.Hour), Status: "予定"},
		{ID: "sch_2", GameID: "game_xyz", StartTime: time.Now().Add(48 * time.Hour), EndTime: time.Now().Add(49 * time.Hour), Status: "予定"},
	}
	return c.JSON(http.StatusOK, mockSchedules)
}

// 3. スケジュールの更新（進捗） [cite: 148]
func updateSchedule(c echo.Context) error {
	id := c.Param("id")
	// リクエストボディから "status": "完了" などを取得

	// ★ TODO: DBの該当スケジュールIDのステータスを更新 ★
	
	return c.JSON(http.StatusOK, map[string]string{
		"message": "スケジュール " + id + " を更新しました",
	})
}

// 4. 継続状態の取得 [cite: 154]
func getMotivationStatus(c echo.Context) error {
	// ★ TODO: DBから現在のポイントやランクを取得 ★

	// 仮のデータを返す
	mockMotivation := Motivation{
		Points: 100,
		Rank:   "ブロンズ",
	}
	return c.JSON(http.StatusOK, mockMotivation)
}

// 4. プレイ結果の報告 [cite: 151]
func reportPlayResult(c echo.Context) error {
	result := new(PlayResult)
	if err := c.Bind(result); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	// ★ TODO: ここでボーナス・ペナルティのロジックを実装 ★
	// result.Result ("completed" or "skipped") に応じてポイントを増減させる

	if result.Result == "completed" {
		// ポイント加算処理
	} else {
		// ペナルティ処理
	}

	// 仮のレスポンス
	return c.JSON(http.StatusOK, map[string]string{
		"message": "プレイ結果 (ID: " + result.ScheduleID + ") を反映しました",
	})
}
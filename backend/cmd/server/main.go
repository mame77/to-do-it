package main

import (
	"database/sql"
	"log"

	"github.com/labstack/echo/v4"
	// _ "github.com/lib/pq" // PostgreSQLドライバ

	// 担当Aのパッケージ
	"TO-DO-IT/internal/calendar"
	"TO-DO-IT/internal/score"

	// 担当Cのパッケージ
	"TO-DO-IT/internal/game" // ← インポートを確認
	// ... (他に必要なパッケージ)
)

func main() {
	// --- DB接続 (仮) ---
	// dsn := "user=... password=... host=... port=... dbname=... sslmode=disable"
	// db, err := sql.Open("postgres", dsn)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()
	var db *sql.DB // (DB接続は仮置き)

	// --- 依存関係の構築 (DI) ---
	// 各担当のリポジトリを初期化
	gameRepo := game.NewRepository(db)     // 担当C
	calendarRepo := calendar.NewRepository(db) // 担当A
	scoreRepo := score.NewRepository(db)       // 担当A
	// ... (taskRepoなど)

	// 各担当のサービスを初期化
	// (担当Aのcalendarサービスは、担当Cのgameリポジトリが必要)
	calendarSvc := calendar.NewService(calendarRepo, gameRepo) // 担当A
	scoreSvc := score.NewService(scoreRepo)                     // 担当A
	
	// ★↓↓↓ 担当Cのサービスを初期化 (コメントアウト解除) ↓↓↓
	gameSvc := game.NewService(gameRepo) 

	// 各担当のハンドラを初期化
	calendarHandler := calendar.NewHandler(calendarSvc) // 担当A
	scoreHandler := score.NewHandler(scoreSvc)          // 担当D
	// ... (gameHandlerなど)

	// --- Echoサーバーのセットアップ ---
	e := echo.New()

	api := e.Group("/api") // /api プレフィックス

	// --- ルート登録 ---
	// 担当Aのルートを登録
	calendarHandler.RegisterRoutes(api)
	scoreHandler.RegisterRoutes(api)

	// 担当Cのルートを登録
	// gameHandler.RegisterRoutes(api)

	// (担当Dのルートを登録 ... )

	// --- サーバー起動 ---
	log.Println("Server starting on :8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
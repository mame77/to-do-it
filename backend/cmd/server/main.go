package main

import (
	// ヘルパー
	"log"
	"time"

	"github.com/mame77/to-do-it/backend/internal/calendar" // 担当A
	"github.com/mame77/to-do-it/backend/internal/db"       // インメモリDB
	"github.com/mame77/to-do-it/backend/internal/game"     // 担当C
	"github.com/mame77/to-do-it/backend/internal/score"    // 担当A

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// --- 1. インメモリDBの初期化 ---
	log.Println("Initializing in-memory database...")
	inMemoryDB := db.NewMemoryDB()
	// (動作確認用のシードデータ)
	seedData(inMemoryDB)

	// --- 2. 依存関係の構築 (DI) ---
	// リポジトリ層 (DBを注入)
	gameRepo := game.NewRepository(inMemoryDB)
	calendarRepo := calendar.NewRepository(inMemoryDB)
	scoreRepo := score.NewRepository(inMemoryDB)

	// サービス層 (リポジトリを注入)
	// (scoreSvcはcalendarSvcで使うので先に作る)
	scoreSvc := score.NewService(scoreRepo)
	gameSvc := game.NewService(gameRepo)
	// (calendarSvcは、担当CのgameRepoと担当AのscoreSvcの両方に依存する)
	calendarSvc := calendar.NewService(calendarRepo, gameRepo, scoreSvc)

	// ハンドラ層 (サービスを注入)
	gameHandler := game.NewHandler(gameSvc)
	calendarHandler := calendar.NewHandler(calendarSvc)
	scoreHandler := score.NewHandler(scoreSvc)

	// --- 3. Echoサーバーのセットアップ ---
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルートグループ
	api := e.Group("/api")

	// --- 4. ルート登録 ---
	// 各担当のハンドラがルートを登録する
	gameHandler.RegisterRoutes(api)     // 担当C
	calendarHandler.RegisterRoutes(api) // 担当A
	scoreHandler.RegisterRoutes(api)    // 担当A

	// --- 5. サーバー起動 ---
	log.Println("Server starting on :8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}

// seedData ... 動作確認用の初期データ
func seedData(mdb *db.MemoryDB) {
	userID := "user_123"

	// 担当C: 所持ゲーム
	game1 := game.Game{ID: "game_001", UserID: userID, Title: "伝説のRPG"}
	game2 := game.Game{ID: "game_002", UserID: userID, Title: "サイバーパンクADV"}
	game3 := game.Game{ID: "game_003", UserID: userID, Title: "宇宙ストラテジー"}
	mdb.Games[game1.ID] = game1
	mdb.Games[game2.ID] = game2
	mdb.Games[game3.ID] = game3

	// 担当A: 固定予定
	// (明日の9時〜17時)
	tomorrow := time.Now().AddDate(0, 0, 1)
	fe1 := calendar.FixedEvent{
		ID:        "fe_001",
		UserID:    userID,
		Title:     "仕事",
		StartTime: time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 9, 0, 0, 0, time.Local),
		EndTime:   time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 17, 0, 0, 0, time.Local),
	}
	mdb.FixedEvents[fe1.ID] = fe1

	// 担当A: モチベーション
	m1 := score.Motivation{
		UserID: userID, Points: 50, Rank: "ブロンズ", Level: 50,
	}
	mdb.Motivations[m1.UserID] = m1
}

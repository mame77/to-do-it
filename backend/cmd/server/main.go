package main

import (
	"database/sql"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3" // SQLiteドライバ

	// 担当Aのパッケージ
	"TO-DO-IT/internal/calendar"
	"TO-DO-IT/internal/score"

	// 担当Cのパッケージ
	"TO-DO-IT/internal/game" // ← インポートを確認
	// ... (他に必要なパッケージ)
)

func main() {
	// --- DB接続 (SQLite) ---
	db, err := sql.Open("sqlite3", "./todo_it.db")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	defer db.Close()

	// テーブルを作成
	if err := initDB(db); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

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
	gameHandler := game.NewHandler(gameSvc)             // 担当C

	// --- Echoサーバーのセットアップ ---
	e := echo.New()

	api := e.Group("/api") // /api プレフィックス

	// --- ルート登録 ---
	// 担当Aのルートを登録
	calendarHandler.RegisterRoutes(api)
	scoreHandler.RegisterRoutes(api)

	// 担当Cのルートを登録
	gameHandler.RegisterRoutes(api)

	// CORS設定を追加
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// --- サーバー起動 ---
	log.Println("Server starting on :8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}

func initDB(db *sql.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS games (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		title TEXT NOT NULL,
		platform TEXT,
		genre TEXT,
		status TEXT DEFAULT 'unstarted',
		release_date DATETIME,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS fixed_events (
		id TEXT PRIMARY KEY,
		user_id TEXT NOT NULL,
		title TEXT NOT NULL,
		start_time DATETIME NOT NULL,
		end_time DATETIME NOT NULL
	);

	CREATE TABLE IF NOT EXISTS schedules (
		id TEXT PRIMARY KEY,
		user_id TEXT NOT NULL,
		game_id TEXT NOT NULL,
		start_time DATETIME NOT NULL,
		end_time DATETIME NOT NULL,
		status TEXT DEFAULT 'pending'
	);

	CREATE TABLE IF NOT EXISTS motivation (
		user_id TEXT PRIMARY KEY,
		points INTEGER DEFAULT 0,
		rank TEXT DEFAULT 'Bronze',
		level INTEGER DEFAULT 1
	);
	`

	_, err := db.Exec(schema)
	return err
}
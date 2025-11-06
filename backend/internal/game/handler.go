package game

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4" // ★GinからEchoに変更
)

// Handler は、game のHTTPリクエスト処理に関するインターフェースです。
type Handler interface {
	RegisterRoutes(apiGroup *echo.Group) // ★引数を *echo.Group に変更
	CreateGame(c echo.Context) error     // ★戻り値に error を追加
	GetGames(c echo.Context) error
	GetGameByID(c echo.Context) error
	UpdateGame(c echo.Context) error
	DeleteGame(c echo.Context) error
}

// handler は Handler インターフェースの具体的な実装です。
type handler struct {
	svc Service
}

// NewHandler は、新しい handler インスタンスを作成します。
func NewHandler(svc Service) Handler {
	return &handler{svc: svc}
}

// RegisterRoutes は、ルーターにエンドポイントを登録します。
func (h *handler) RegisterRoutes(apiGroup *echo.Group) { // ★引数を *echo.Group に変更
	gameRoutes := apiGroup.Group("/games") // /api/games がベースになる
	{
		gameRoutes.POST("", h.CreateGame)       // POST /api/games
		gameRoutes.GET("", h.GetGames)          // GET /api/games
		gameRoutes.GET("/:id", h.GetGameByID)   // GET /api/games/:id
		gameRoutes.PUT("/:id", h.UpdateGame)    // PUT /api/games/:id
		gameRoutes.DELETE("/:id", h.DeleteGame) // DELETE /api/games/:id
	}
}

// --- 個々のハンドラ実装 (Echo形式) ---

// getIDParam は URL から :id をそのまま取得するヘルパー関数
func getIDParam(c echo.Context) (string, error) {
	id := c.Param("id")
	if id == "" {
		return "", echo.NewHTTPError(http.StatusBadRequest, "Missing ID parameter")
	}
	return id, nil
}

// CreateGame は新しいゲームを作成します (POST /api/games)
func (h *handler) CreateGame(c echo.Context) error {
	var req CreateGameRequest

	// 1. リクエストボディ(JSON)を req 構造体にバインド
	// Echoでは c.Bind() を使う
	if err := c.Bind(&req); err != nil {
		log.Printf("Handler: Failed to bind JSON: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body: " + err.Error()})
	}

	// 2. サービスを呼び出す
	game, err := h.svc.CreateGame(&req)
	if err != nil {
		log.Printf("Handler: Error creating game: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create game"})
	}

	// 3. 成功レスポンス（作成されたリソース）を返す
	return c.JSON(http.StatusCreated, game)
}

// GetGames は（テストユーザーの）ゲーム一覧を取得します (GET /api/games)
func (h *handler) GetGames(c echo.Context) error {
	games, err := h.svc.GetGames()
	if err != nil {
		log.Printf("Handler: Error getting games: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get games"})
	}
	return c.JSON(http.StatusOK, games)
}

// GetGameByID は ID でゲームを1件取得します (GET /api/games/:id)
func (h *handler) GetGameByID(c echo.Context) error {
	id, err := getIDParam(c)
	if err != nil {
		return err // getIDParamがHTTPErrorを返しているのでそのまま返す
	}

	game, err := h.svc.GetGame(id)
	if err != nil {
		log.Printf("Handler: Error getting game by ID: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get game"})
	}

	if game == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Game not found"})
	}

	return c.JSON(http.StatusOK, game)
}

// UpdateGame はゲーム情報を更新します (PUT /api/games/:id)
func (h *handler) UpdateGame(c echo.Context) error {
	id, err := getIDParam(c)
	if err != nil {
		return err
	}

	var req UpdateGameRequest
	if err := c.Bind(&req); err != nil {
		log.Printf("Handler: Failed to bind JSON for update: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body: " + err.Error()})
	}

	game, err := h.svc.UpdateGame(id, &req)
	if err != nil {
		log.Printf("Handler: Error updating game: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update game"})
	}

	if game == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Game not found to update"})
	}

	return c.JSON(http.StatusOK, game)
}

// DeleteGame は ID を指定してゲームを削除します (DELETE /api/games/:id)
func (h *handler) DeleteGame(c echo.Context) error {
	id, err := getIDParam(c)
	if err != nil {
		return err
	}

	err = h.svc.DeleteGame(id)
	if err != nil {
		log.Printf("Handler: Error deleting game: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete game"})
	}

	return c.NoContent(http.StatusNoContent) // 中身なし
}

package game

import (
	"TO-DO-IT/internal/app" // (main.goで定義するヘルパー)
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{service: s}
}

// RegisterRoutes ... EchoルーターにAPIエンドポイントを登録
func (h *Handler) RegisterRoutes(api *echo.Group) {
	gameApi := api.Group("/games") // /api/games
	{
		gameApi.POST("", h.handleCreateGame)
		gameApi.GET("/pending", h.handleGetPendingGames)
	}
}

func (h *Handler) handleCreateGame(c echo.Context) error {
	// リクエストボディのパース
	var reqBody struct {
		Title string `json:"title"`
		Genre string `json:"genre"`
	}
	if err := c.Bind(&reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request body")
	}

	// 認証済みUserIDを取得 (今回はヘルパーから固定値)
	userID := app.GetUserIDFromContext(c)

	game, err := h.service.CreateGame(userID, reqBody.Title, reqBody.Genre)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, game)
}

func (h *Handler) handleGetPendingGames(c echo.Context) error {
	userID := app.GetUserIDFromContext(c)
	games, err := h.service.GetPendingGames(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, games)
}

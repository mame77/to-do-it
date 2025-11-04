package calendar

import (
	"net/http"
	"time"

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
	calApi := api.Group("/calendar") // /api/calendar
	{
		// 自動生成 [cite: 71]
		calApi.POST("/generate", h.handleGenerateSchedule)

		// スケジュール [cite: 72-73]
		calApi.GET("/schedule", h.handleGetSchedules)
		calApi.PUT("/schedule/:id", h.handleUpdateScheduleStatus)

		// 固定予定 [cite: 81-82]
		calApi.POST("/fixed-events", h.handleCreateFixedEvent)
		calApi.GET("/fixed-events", h.handleGetFixedEvents)
	}
}

// --- ハンドラの実装 ---

func (h *Handler) handleGenerateSchedule(c echo.Context) error {
	// TODO: JWTなどからUserIDを取得
	userID := "user_123" // 仮

	schedules, err := h.service.GenerateSchedule(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, schedules)
}

func (h *Handler) handleGetSchedules(c echo.Context) error {
	userID := "user_123" // 仮
	// TODO: クエリパラメータから期間を取得
	start := time.Now()
	end := time.Now().Add(7 * 24 * time.Hour)

	schedules, err := h.service.GetSchedules(userID, start, end)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, schedules)
}

func (h *Handler) handleUpdateScheduleStatus(c echo.Context) error {
	scheduleID := c.Param("id")

	// リクエストボディから新しいステータスを取得
	var reqBody struct {
		Status string `json:"status"`
	}
	if err := c.Bind(&reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request body")
	}

	if err := h.service.UpdateScheduleStatus(scheduleID, reqBody.Status); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "status updated"})
}

func (h *Handler) handleCreateFixedEvent(c echo.Context) error {
	var event FixedEvent
	if err := c.Bind(&event); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request body")
	}
	event.UserID = "user_123" // 仮

	if err := h.service.CreateFixedEvent(&event); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, event)
}

func (h *Handler) handleGetFixedEvents(c echo.Context) error {
	userID := "user_123" // 仮
	start := time.Now()
	end := time.Now().Add(7 * 24 * time.Hour)

	events, err := h.service.GetFixedEvents(userID, start, end)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, events)
}

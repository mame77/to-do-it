package calendar

import (
	"TO-DO-IT/internal/app" // (main.goで定義するヘルパー)
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
		// スケジュール (自動生成)
		[cite_start]calApi.POST("/generate", h.handleGenerateSchedule)        // [cite: 147]
		[cite_start]calApi.GET("/schedule", h.handleGetSchedules)             // [cite: 147]
		[cite_start]calApi.PUT("/schedule/:id", h.handleUpdateScheduleStatus) // [cite: 148]

		// 固定予定 (手動)
		[cite_start]calApi.POST("/fixed-events", h.handleCreateFixedEvent) // [cite: 156]
		[cite_start]calApi.GET("/fixed-events", h.handleGetFixedEvents)    // [cite: 157]
	}
}

// --- Schedule Handlers ---

func (h *Handler) handleGenerateSchedule(c echo.Context) error {
	userID := app.GetUserIDFromContext(c)
	schedules, err := h.service.GenerateSchedule(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, schedules)
}

func (h *Handler) handleGetSchedules(c echo.Context) error {
	userID := app.GetUserIDFromContext(c)
	schedules, err := h.service.GetSchedules(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, schedules)
}

func (h *Handler) handleUpdateScheduleStatus(c echo.Context) error {
	scheduleID := c.Param("id")

	// リクエストボディから新しいステータスを取得
	var reqBody struct {
		Status string `json:"status"` // "completed" or "skipped"
	}
	if err := c.Bind(&reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request body")
	}

	if err := h.service.UpdateScheduleStatus(scheduleID, reqBody.Status); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "status updated"})
}

// --- FixedEvent Handlers ---

func (h *Handler) handleCreateFixedEvent(c echo.Context) error {
	var reqBody struct {
		Title string `json:"title"`
		Start string `json:"start_time"` // "2025-11-10T09:00:00Z" (RFC3339)
		End   string `json:"end_time"`   // "2025-11-10T17:00:00Z" (RFC3339)
	}
	if err := c.Bind(&reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request body")
	}

	startTime, err1 := time.Parse(time.RFC3339, reqBody.Start)
	endTime, err2 := time.Parse(time.RFC3339, reqBody.End)
	if err1 != nil || err2 != nil {
		return c.JSON(http.StatusBadRequest, "invalid time format, use RFC3339")
	}

	userID := app.GetUserIDFromContext(c)
	event, err := h.service.CreateFixedEvent(userID, reqBody.Title, startTime, endTime)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, event)
}

func (h *Handler) handleGetFixedEvents(c echo.Context) error {
	userID := app.GetUserIDFromContext(c)
	events, err := h.service.GetFixedEvents(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, events)
}

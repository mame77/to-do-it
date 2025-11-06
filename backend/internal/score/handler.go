package score

import (
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
	scoreApi := api.Group("/motivation") // /api/motivation
	{
		scoreApi.GET("", h.handleGetMotivation)        // [cite: 79]
		scoreApi.POST("/result", h.handleReportResult) // [cite: 76]
	}
}

func (h *Handler) handleGetMotivation(c echo.Context) error {
	userID := "user_123" // 仮
	motivation, err := h.service.GetMotivation(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, motivation)
}

func (h *Handler) handleReportResult(c echo.Context) error {
	userID := "user_123" // 仮
	var result PlayResult
	if err := c.Bind(&result); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request body")
	}

	motivation, err := h.service.ReportPlayResult(userID, result)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, motivation)
}

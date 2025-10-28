package calendar

import (
	"encoding/json"
	"net/http"
)

// Handler はHTTPリクエストを処理します。
type Handler struct {
	service *Service
}

// NewHandler はHandlerを初期化します。
func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// GenerateSchedulesHandler はスケジュールの自動生成エンドポイントです。
// POST /api/v1/calendar/generate
func (h *Handler) GenerateSchedulesHandler(w http.ResponseWriter, r *http.Request) {
	var req GenerateScheduleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	schedules, err := h.service.GenerateSchedules(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(schedules); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// ScheduleActionHandler はスケジュールの完了・スキップ操作のエンドポイントです。
// POST /api/v1/calendar/action
func (h *Handler) ScheduleActionHandler(w http.ResponseWriter, r *http.Request) {
	var req ScheduleActionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateScheduleAction(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// --- 💡 ルーティング例（メインアプリケーションでこれらを登録する） ---

/*
func RegisterRoutes(router *http.ServeMux, handler *Handler) {
	router.HandleFunc("/api/v1/calendar/generate", handler.GenerateSchedulesHandler)
	router.HandleFunc("/api/v1/calendar/action", handler.ScheduleActionHandler)
}
*/
package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/aldihorn/pact-cfp-demo/apps/api/internal/service"
)

type TodoHandler struct {
	service service.TodoService
}

func NewTodoHandler(svc service.TodoService) *TodoHandler {
	return &TodoHandler{service: svc}
}

type createTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type patchTodoRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Completed   *bool   `json:"completed"`
}

func (h *TodoHandler) ListTodos(w http.ResponseWriter, r *http.Request) {
	var completed *bool
	queryValue := r.URL.Query().Get("completed")
	if queryValue != "" {
		parsed, err := strconv.ParseBool(queryValue)
		if err != nil {
			writeJSONError(w, http.StatusBadRequest, "invalid completed query value")
			return
		}
		completed = &parsed
	}

	todos, err := h.service.ListTodos(r.Context(), completed)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "failed to list todos")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"data": todos})
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var req createTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	todo, err := h.service.CreateTodo(r.Context(), service.CreateTodoInput{
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		h.mapServiceError(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, map[string]any{"data": todo})
}

func (h *TodoHandler) PatchTodo(w http.ResponseWriter, r *http.Request) {
	id, err := parseTodoID(r)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	var req patchTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	todo, err := h.service.UpdateTodo(r.Context(), id, service.UpdateTodoInput{
		Title:       req.Title,
		Description: req.Description,
		Completed:   req.Completed,
	})
	if err != nil {
		h.mapServiceError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"data": todo})
}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := parseTodoID(r)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.service.DeleteTodo(r.Context(), id); err != nil {
		h.mapServiceError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *TodoHandler) mapServiceError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, service.ErrInvalidTitle):
		writeJSONError(w, http.StatusBadRequest, err.Error())
	case errors.Is(err, service.ErrTodoNotFound):
		writeJSONError(w, http.StatusNotFound, err.Error())
	default:
		writeJSONError(w, http.StatusInternalServerError, "internal server error")
	}
}

func parseTodoID(r *http.Request) (uint, error) {
	idValue := chi.URLParam(r, "id")
	parsed, err := strconv.ParseUint(idValue, 10, 64)
	if err != nil {
		return 0, errors.New("invalid todo id")
	}
	return uint(parsed), nil
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func writeJSONError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]any{"error": message})
}

package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/aldihorn/pact-cfp-demo/apps/api/internal/api/handler"
	appmw "github.com/aldihorn/pact-cfp-demo/apps/api/internal/api/middleware"
	"github.com/aldihorn/pact-cfp-demo/apps/api/internal/config"
	"github.com/aldihorn/pact-cfp-demo/apps/api/internal/observability"
	"github.com/aldihorn/pact-cfp-demo/apps/api/internal/service"
)

func NewRouter(todoService service.TodoService, _ config.Config) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(appmw.Logging)

	todoHandler := handler.NewTodoHandler(todoService)

	r.Get("/healthz", handler.Healthz)
	r.Get("/readyz", handler.Readyz)
	r.Handle("/metrics", observability.MetricsHandler())

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/todos", todoHandler.ListTodos)
		r.Post("/todos", todoHandler.CreateTodo)
		r.Patch("/todos/{id}", todoHandler.PatchTodo)
		r.Delete("/todos/{id}", todoHandler.DeleteTodo)
	})

	return r
}

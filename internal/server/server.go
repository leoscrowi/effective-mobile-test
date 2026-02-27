package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/leoscrowi/effective-mobile-test/docs"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"
)

type RouteSetup interface {
	SetupRoutes(r chi.Router)
}

type Server struct {
	Router      chi.Router
	Controllers []RouteSetup
}

func NewServer(db *gorm.DB) *Server {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	return &Server{
		Router:      r,
		Controllers: GetControllers(db),
	}
}

func (s *Server) SetupRoutes() {
	for _, controller := range s.Controllers {
		controller.SetupRoutes(s.Router)
	}
}

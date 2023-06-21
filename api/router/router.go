package router

import (
	api "test-go-server/api/api"
	"test-go-server/logger"

	middlewares "test-go-server/middlewares"
	resource "test-go-server/pkg/resource"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func New(l logger.Logger, rs resource.Resources) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Route("/public", func(r chi.Router) {
		API := api.NewAPI(rs, l)
		r.Get("/", API.List)
		r.Post("/{resource}", API.Create)
		r.Get("/{resource}", API.Read)
		r.Put("/{resource}", API.Update)
		r.Delete("/{resource}", API.Delete)
	})

	r.Route("/private", func(r chi.Router) {
		API := api.NewAPI(rs, l)
		auth := middlewares.AuthMiddleware{AuthToken: "123"}
		r.Use(auth.Run)
		r.Get("/", API.List)
		r.Post("/{resource}", API.Create)
		r.Get("/{resource}", API.Read)
		r.Put("/{resource}", API.Update)
		r.Delete("/{resource}", API.Delete)
	})

	return r
}

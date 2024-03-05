package routing

import (
	"github.com/daniskazan/url-shortener/internal/delivery/rest/users"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func configureMiddleware(router chi.Router) {
	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)
}

func ConfigureRouting() *chi.Mux {
	apiRouter := chi.NewRouter()
	configureMiddleware(apiRouter)
	apiRouter.Route("/users", loadUserRoutes)

	return apiRouter
}

func loadUserRoutes(router chi.Router) {
	userHandler := users.UserHandler{}

	router.Get("/{id}", userHandler.GetUserByID)
	router.Get("/", userHandler.GetUserList)
	router.Post("/", userHandler.CreateUser)
	router.Patch("/{id}", userHandler.UpdateUser)
	router.Delete("/{id}", userHandler.DeleteUserByID)
}

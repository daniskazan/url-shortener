package application

import (
	"github.com/daniskazan/url-shortener/configs"
	user_handler "github.com/daniskazan/url-shortener/internal/delivery/rest/users"
	"github.com/daniskazan/url-shortener/internal/repository/users"
	"github.com/go-chi/chi/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

type Application struct {
	db     gorm.DB
	router chi.Router
	config configs.ApplicationConfig
}

func NewApplication() *Application {
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	app := &Application{db: *db, router: chi.NewRouter(), config: configs.ApplicationConfig{Addr: ":8000"}}
	return app
}

func (a *Application) ConfigureRouting() {
	a.router.Route("/users", a.ConfigureUserRouting)
}

func (a *Application) ConfigureUserRouting(router chi.Router) {
	userRepo := users.UserRepository{DB: &a.db}
	userHandler := user_handler.UserHandler{Repo: &userRepo}

	router.Get("/{id}", userHandler.GetUserByID)
	router.Get("/", userHandler.GetUserList)
	router.Post("/", userHandler.CreateUser)
	router.Patch("/{id}", userHandler.UpdateUser)
	router.Delete("/{id}", userHandler.DeleteUserByID)

}
func (a *Application) Start() error {
	err := http.ListenAndServe(a.config.Addr, a.router)
	if err != nil {
		panic(err)
	}
	return nil
}

func (a *Application) ShutDown() error {
	return nil
}

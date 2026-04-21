package main

import (
	"database/sql"
	"net/http"
	"tic_tac_toe/internal/front"
	"tic_tac_toe/internal/handler"
	"tic_tac_toe/internal/repository"
	"tic_tac_toe/internal/service"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
)

const BaseUrl = "localhost:8080"

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5433 user=gamer password=gaming_pass dbname=games_db sslmode=disable")
	if err != nil {
		return
	}
	defer db.Close()
	fx.New(
		fx.Supply(db),
		fx.Provide(repository.NewGameRepository, service.NewGameService, handler.NewGameHandler, chi.NewRouter,
			front.NewGameApp, front.NewStartWindow),
		fx.Invoke(func(rout *chi.Mux, handlers *handler.GameHandler) {
			rout.Post("/games/{id}", handlers.PostHandler)
			rout.Get("/games/{id}", handlers.GetHandler)
			rout.Delete("/games/{id}", handlers.DeleteHandler)
			rout.Get("/games", handlers.GetAllHandler)
		}),
		fx.Invoke(func(shutdowner fx.Shutdowner, lc fx.Lifecycle, app *front.StartWindow, rout *chi.Mux, service *service.GameService) {
			server := http.Server{Addr: BaseUrl, Handler: rout}
			go server.ListenAndServe()
			lc.Append(fx.Hook{OnStop: server.Shutdown})
			app.ShowStartWindow(service)
			shutdowner.Shutdown()
		}),
	).Run()
}

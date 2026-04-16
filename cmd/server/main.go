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
)

const BaseUrl = "localhost:8080"

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=Dimon007 dbname=Games sslmode=disable")
	if err != nil {
		return
	}
	repository := repository.NewGameRepository(db)
	service := service.NewGameService(repository)
	handlers := handler.NewGameHandler(service)
	rout := chi.NewRouter()
	rout.Post("/games/{id}", handlers.PostHandler)
	rout.Get("/games/{id}", handlers.GetHandler)
	rout.Delete("/games/{id}", handlers.DeleteHandler)
	rout.Get("/games", handlers.GetAllHandler)
	go http.ListenAndServe(BaseUrl, rout)
	a := front.NewGameApp()
	app := front.NewStartWindow(a)
	app.ShowStartWindow(service)
}

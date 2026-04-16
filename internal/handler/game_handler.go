package handler

import (
	"encoding/json"
	"net/http"
	"tic_tac_toe/internal/logic"
	"tic_tac_toe/internal/service"

	"github.com/go-chi/chi"
	"github.com/shabbyrobe/go-num"
)

type GameHandler struct {
	service *service.GameService
}

func NewGameHandler(s *service.GameService) *GameHandler {
	return &GameHandler{service: s}
}

func (g *GameHandler) PostHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idU128, _, err := num.U128FromString(id)
	if err != nil {
		http.Error(w, "Incorrect id", http.StatusBadRequest)
		return
	}
	if check, err := g.service.Db.CheckGame(idU128); err != nil || check {
		return
	}
	game := logic.NewGameLogic(idU128, 1)
	g.service.Db.CreateGame(game)
	go g.service.M.Store(idU128, &game)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (g *GameHandler) GetHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idU128, _, err := num.U128FromString(id)
	if err != nil {
		http.Error(w, "Incorrect id", http.StatusBadRequest)
		return
	}
	game, err := g.service.Db.GetGame(idU128)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(game)
}

func (g *GameHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idU128, _, err := num.U128FromString(id)
	if err != nil {
		http.Error(w, "Incorrect id", http.StatusBadRequest)
		return
	}
	check, err := g.service.Db.CheckGame(idU128)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if check == false {
		http.Error(w, "Not found id", http.StatusNotFound)
		return
	}
	err = g.service.Db.DeleteGame(idU128)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, ok := g.service.M.Load(idU128); ok {
		go g.service.M.Delete(idU128)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (g *GameHandler) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	games, err := g.service.Db.GetAllGames()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(games)
}

package service

import (
	"sync"
	"tic_tac_toe/internal/repository"
)

type GameService struct {
	Db *repository.GameRepository
	M  sync.Map
}

func NewGameService(Db *repository.GameRepository) *GameService {
	return &GameService{Db: Db}
}

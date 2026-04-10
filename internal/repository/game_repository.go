package repository

import (
	"database/sql"
	"encoding/json"
	"errors"
	"tic_tac_toe/internal/logic"
	"tic_tac_toe/internal/models"

	"github.com/shabbyrobe/go-num"
)

type GameRepository struct {
	db *sql.DB
}

func NewGameRepository(db *sql.DB) *GameRepository {
	return &GameRepository{db: db}
}

func (db *GameRepository) CreateGame(g *logic.GameLogic) error {
	jsonData, err := json.Marshal(g.Board.Condition)
	if err != nil {
		return errors.New("Fail tranformation to JSON")
	}
	query := "INSERT INTO GAMES(id, condition, player) VALUES($1, $2, $3)"
	_, err = db.db.Exec(query, g.UUid.String(), string(jsonData), g.Player)
	if err != nil {
		return errors.New("Fail INSERT")
	}
	return nil
}

func (db *GameRepository) DeleteGame(idU128 num.U128) error {
	query := "DELETE FROM GAMES WHERE id = $1"
	_, err := db.db.Exec(query, idU128.String())
	if err != nil {
		return errors.New("Fail to DELETE")
	}
	return nil
}

func (db *GameRepository) UpdateGame(g *logic.GameLogic) error {
	jsonData, err := json.Marshal(g.Board.Condition)
	if err != nil {
		return errors.New("Fail to Marshall")
	}
	query := "UPDATE GAMES SET condition = $1 WHERE id = $2"
	_, err = db.db.Exec(query, string(jsonData), g.UUid.String())
	if err != nil {
		return errors.New("Fail to Update")
	}
	return nil
}

func (db *GameRepository) GetGame(searchId num.U128) (*logic.GameLogic, error) {
	var id string
	var condition string
	var player int8
	var gameCondition [3][3]int8
	query := "SELECT id, condition, player FROM games WHERE id = $1"
	err := db.db.QueryRow(query, searchId.String()).Scan(&id, &condition, &player)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(condition), &gameCondition)
	if err != nil {
		return nil, errors.New("Fail to Marshall")
	}
	dbUUid, _, _ := num.U128FromString(id)
	return &logic.GameLogic{UUid: dbUUid, Board: models.Board{Condition: gameCondition}, Player: player}, nil
}

func (db *GameRepository) CheckGame(idU128 num.U128) (bool, error) {
	query := "SELECT 1 FROM games WHERE id = $1"
	var row string
	err := db.db.QueryRow(query, idU128.String()).Scan(&row)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, errors.New("Incorrect type")
	}
	return row != "", nil
}

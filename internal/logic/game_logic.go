package logic

import (
	"tic_tac_toe/internal/models"

	"github.com/shabbyrobe/go-num"
)

type GameLogic struct {
	Player int8         `json:"player"`
	UUid   num.U128     `json:"uuid"`
	Board  models.Board `json:"board"`
}

func NewGameLogic(UUid num.U128, player int8) *GameLogic {
	return &GameLogic{UUid: UUid, Player: player}
}

func (g *GameLogic) PlayerStep(row, column uint8) {
	g.Board.Condition[row][column] = g.Player
}

func (g *GameLogic) EndGame() bool {
	return CheckWinnerBot(g.Board.Condition) || CheckWinnerPlayer(g.Board.Condition) || fullBoard(g.Board.Condition)
}
func CheckWinnerPlayer(board [3][3]int8) bool {
	for i := 0; i < 3; i++ {
		if (board[i][0]+board[i][1]+board[i][2] == 3) ||
			(board[0][i]+board[1][i]+board[2][i] == 3) {
			return true
		}
	}
	if (board[0][0]+board[1][1]+board[2][2] == 3) ||
		(board[0][2]+board[1][1]+board[2][0] == 3) {
		return true
	}
	return false
}

func CheckWinnerBot(board [3][3]int8) bool {
	for i := 0; i < 3; i++ {
		if (board[i][0]+board[i][1]+board[i][2] == -3) ||
			(board[0][i]+board[1][i]+board[2][i] == -3) {
			return true
		}
	}
	if (board[0][0]+board[1][1]+board[2][2] == -3) ||
		(board[0][2]+board[1][1]+board[2][0] == -3) {
		return true
	}
	return false
}

func fullBoard(board [3][3]int8) bool {
	for _, i := range board {
		for _, j := range i {
			if j == 0 {
				return false
			}
		}
	}
	return true
}

func (g *GameLogic) MachineStep() {
	copy := g.Board.Condition
	max := struct {
		max    int
		row    int
		column int
	}{2, 0, 0}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if copy[i][j] == 0 {
				copy[i][j] = -1
				result := minimax(copy, false)
				copy[i][j] = 0
				if result < max.max {
					max = struct {
						max    int
						row    int
						column int
					}{result, i, j}
				}
			}
		}
	}
	g.Board.Condition[max.row][max.column] = -1
}

func minimax(board [3][3]int8, bot bool) int {
	if CheckWinnerBot(board) {
		return -1
	} else if CheckWinnerPlayer(board) {
		return 1
	} else if fullBoard(board) {
		return 0
	}
	var score int
	if bot == true {
		score = 1
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[i][j] == 0 {
					board[i][j] = -1
					result := minimax(board, !bot)
					board[i][j] = 0
					if result < score {
						score = result
					}
				}
			}
		}
	} else {
		score = -1
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[i][j] == 0 {
					board[i][j] = 1
					result := minimax(board, !bot)
					board[i][j] = 0
					if result > score {
						score = result
					}
				}
			}
		}
	}
	return score
}

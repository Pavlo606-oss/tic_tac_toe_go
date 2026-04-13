package front

import "tic_tac_toe/internal/logic"

func buttonString(game *logic.GameLogic, i, j int8) string {
	switch game.Board.Condition[i][j] {
	case 1:
		return "X"
	case -1:
		return "O"
	}
	return ""
}

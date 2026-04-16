package front

import (
	"net/http"
	"tic_tac_toe/internal/logic"
	"tic_tac_toe/internal/service"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/shabbyrobe/go-num"
)

func ShowEndDialog(window *fyne.Window, idU128 num.U128, gs *service.GameService, pgw *PlayingGameWindow, game *logic.GameLogic) {
	dialog.ShowConfirm("Игра закончена!", WinnerName(game)+"\nХотите начать игру с тем же id заново?", func(res bool) {
		if res {
			req, _ := http.NewRequest(http.MethodDelete, "http://localhost:8080/games/"+idU128.String(), nil)
			http.DefaultClient.Do(req)
			_, err := http.Post("http://localhost:8080/games/"+idU128.String(), "application/json", nil)
			if err != nil {
				return
			}
			go gs.M.Store(idU128, logic.NewGameLogic(idU128, 1))
			gs.Db.CreateGame(logic.NewGameLogic(idU128, 1))
			playingWindow := NewPlayingGameWindow(pgw.app)
			playingWindow.ShowNewPlayingGameWindow(idU128, gs)
		}
		(*window).Close()
	}, *window)
}

func ShowExistGameDialog(ngw *NewGameWindow) {
	dialog.ShowInformation("Крестики-нолики", "Такая игра уже существует", ngw.window)
}

func ShowNotExistGameDialog(cgw *ChoiceGameWindow) {
	dialog.ShowInformation("Крестики-нолики", "Такой игры не существует", cgw.window)
}

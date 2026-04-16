package front

import (
	"encoding/json"
	"net/http"
	"tic_tac_toe/internal/logic"

	"fyne.io/fyne/v2"
)

type DeleteWindow struct {
	app    *GameApp
	window fyne.Window
}

func NewDeleteWindow(app *GameApp) *DeleteWindow {
	return &DeleteWindow{app: app}
}

func (dw *DeleteWindow) ShowDeleteWindow() {
	SettingDeleteWindow(dw)
	listGames := make([]logic.GameLogic, 0)
	resp, _ := http.Get("http://localhost:8080/games")
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&listGames)
	SetListGamesOnDeleteWindow(dw, listGames)
	dw.window.Show()
}

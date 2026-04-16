package front

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tic_tac_toe/internal/logic"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/shabbyrobe/go-num"
)

func SettingsPlayingWindow(pgw *PlayingGameWindow, idU128 num.U128) {
	pgw.window = pgw.app.app.NewWindow(fmt.Sprintf("Игра %s", idU128.String()))
	pgw.window.SetFixedSize(true)
	pgw.window.Resize(fyne.Size{Width: 375, Height: 375})
}

func SettingDeleteWindow(dw *DeleteWindow) {
	dw.window = dw.app.app.NewWindow("Просмотр игр")
	dw.window.Resize(fyne.Size{Width: 375, Height: 375})
}

func SetListGamesOnDeleteWindow(dw *DeleteWindow, listGames []logic.GameLogic) {
	container := container.NewVBox()
	infoLabel := widget.NewLabel("ID ИГРЫ        СОСТОЯНИЕ ПОЛЯ")
	container.Add(infoLabel)
	for i, game := range listGames {
		conditionString, _ := json.Marshal(game.Board.Condition)
		label := widget.NewLabel(game.UUid.String() + "		  " + string(conditionString))
		button := widget.NewButton("Удалить", func() {
			index := i
			idU128String := game.UUid.String()
			req, _ := http.NewRequest(http.MethodDelete, "http://localhost:8080/games/"+idU128String, nil)
			http.DefaultClient.Do(req)
			listGames = append(listGames[:index], listGames[index+1:]...)
			SetListGamesOnDeleteWindow(dw, listGames)
		})
		container.Add(label)
		container.Add(button)
	}
	dw.window.SetContent(container)
}

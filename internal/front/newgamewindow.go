package front

import (
	"net/http"
	"tic_tac_toe/internal/logic"
	"tic_tac_toe/internal/service"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/shabbyrobe/go-num"
)

type NewGameWindow struct {
	app    *GameApp
	window fyne.Window
}

func NewNewGameWindow(app *GameApp) *NewGameWindow {
	return &NewGameWindow{app: app}
}

func (ngw *NewGameWindow) ShowNewGameWindow(gs *service.GameService) {
	ngw.window = ngw.app.app.NewWindow("Крестики-нолики")
	ngw.window.Resize(fyne.Size{Width: 400, Height: 400})
	ngw.window.SetFixedSize(true)
	label := widget.NewLabel("Пожалуйста, введите id для новой игры")
	entry := widget.NewEntry()
	entry.Resize(fyne.NewSize(200, 50))
	entry.Move(fyne.NewPos(100, 50))
	enterButton := widget.NewButton("Готово", func() {
		idU128, _, _ := num.U128FromString(entry.Text)
		if check, err := gs.Db.CheckGame(idU128); err != nil || check {
			ShowExistGameDialog(ngw)
			return
		}
		_, err := http.Post("http://localhost:8080/games/"+entry.Text, "application/json", nil)
		if err != nil {
			return
		}
		go gs.M.Store(idU128, logic.NewGameLogic(idU128, 1))
		playingWindow := NewPlayingGameWindow(ngw.app)
		playingWindow.ShowNewPlayingGameWindow(idU128, gs)
	})
	enterButton.Move(fyne.NewPos(150, 150))
	enterButton.Resize(fyne.NewSize(100, 50))
	enterButton.Resize(fyne.NewSize(100, 50))
	label.Move(fyne.NewPos(50, 10))
	ngw.window.SetContent(container.NewWithoutLayout(label, entry, enterButton))
	ngw.window.Show()
}

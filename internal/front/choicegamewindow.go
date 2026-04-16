package front

import (
	"encoding/json"
	"net/http"
	"tic_tac_toe/internal/logic"
	"tic_tac_toe/internal/service"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/shabbyrobe/go-num"
)

type ChoiceGameWindow struct {
	app    *GameApp
	window fyne.Window
}

func NewChoiceGameWindow(app *GameApp) *ChoiceGameWindow {
	return &ChoiceGameWindow{app: app}
}

func (cgw *ChoiceGameWindow) ShowChoiceGameWindow(gs *service.GameService) {
	cgw.window = cgw.app.app.NewWindow("Крестики-нолики")
	label := widget.NewLabel("Пожалуйста, введите id игры, которую хотите выбрать")
	entry := widget.NewEntry()
	entry.Resize(fyne.NewSize(200, 50))
	entry.Move(fyne.NewPos(100, 50))
	enterButton := widget.NewButton("Готово", func() {
		idU128, _, _ := num.U128FromString(entry.Text)
		if _, ok := gs.M.Load(idU128); !ok {
			resp, err := http.Get("http://localhost:8080/games/" + entry.Text)
			defer func() {
				resp.Body.Close()
			}()
			if err != nil {
				return
			}
			game := logic.GameLogic{}
			if err := json.NewDecoder(resp.Body).Decode(&game); err != nil {
				ShowNotExistGameDialog(cgw)
				return
			}
			go gs.M.Store(idU128, &game)
		}
		playingWindow := NewPlayingGameWindow(cgw.app)
		playingWindow.ShowNewPlayingGameWindow(idU128, gs)
	})
	enterButton.Move(fyne.NewPos(150, 150))
	enterButton.Resize(fyne.NewSize(100, 50))
	enterButton.Resize(fyne.NewSize(100, 50))
	label.Move(fyne.NewPos(5, 10))
	cgw.window.SetContent(container.NewWithoutLayout(label, entry, enterButton))
	cgw.window.Resize(fyne.Size{Width: 400, Height: 400})
	cgw.window.SetFixedSize(true)
	cgw.window.Show()
}

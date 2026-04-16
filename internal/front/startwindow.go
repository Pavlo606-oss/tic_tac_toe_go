package front

import (
	"tic_tac_toe/internal/service"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type StartWindow struct {
	app    *GameApp
	window fyne.Window
}

func NewStartWindow(app *GameApp) *StartWindow {
	return &StartWindow{app: app}
}

func (sw *StartWindow) ShowStartWindow(gs *service.GameService) {
	sw.window = sw.app.app.NewWindow("Крестики-нолики")
	hello := widget.NewLabel("Выберите действие")
	newGameButton := widget.NewButton("Новая игра", func() {
		newGameWindow := NewNewGameWindow(sw.app)
		newGameWindow.ShowNewGameWindow(gs)
	})
	choiceGameButton := widget.NewButton("Выбрать игру по ID", func() {
		choiceGameWindow := NewChoiceGameWindow(sw.app)
		choiceGameWindow.ShowChoiceGameWindow(gs)
	})
	checkGamesButton := widget.NewButton("Посмотреть созданные игры", func() {
		checkGamesWindow := NewDeleteWindow(sw.app)
		checkGamesWindow.ShowDeleteWindow()
	})
	newGameButton.Move(fyne.NewPos(50, 150))
	choiceGameButton.Move(fyne.NewPos(200, 150))
	checkGamesButton.Move(fyne.NewPos(75, 250))
	newGameButton.Resize(fyne.NewSize(100, 50))
	choiceGameButton.Resize(fyne.NewSize(150, 50))
	checkGamesButton.Resize(fyne.NewSize(250, 50))
	hello.Move(fyne.NewPos(120, 50))
	sw.window.SetContent(container.NewWithoutLayout(hello, newGameButton, choiceGameButton, checkGamesButton))
	sw.window.Resize(fyne.Size{Width: 400, Height: 400})
	sw.window.SetFixedSize(true)
	sw.window.ShowAndRun()
}

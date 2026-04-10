package front

import (
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

func (sw *StartWindow) ShowStartWindow() {
	sw.window = sw.app.app.NewWindow("Крестики-нолики")
	hello := widget.NewLabel("Выберите действие")
	newGameButton := widget.NewButton("Новая игра", func() {
		newGameWindow := NewNewGameWindow(sw.app)
		newGameWindow.ShowNewGameWindow()
	})
	choiceGameButton := widget.NewButton("Выбрать игру по ID", func() {
		choiceGameWindow := NewChoiceGameWindow(sw.app)
		choiceGameWindow.ShowChoiceGameWindow()
	})
	newGameButton.Move(fyne.NewPos(50, 150))
	choiceGameButton.Move(fyne.NewPos(250, 150))
	newGameButton.Resize(fyne.NewSize(100, 50))
	choiceGameButton.Resize(fyne.NewSize(100, 50))
	hello.Move(fyne.NewPos(120, 50))
	sw.window.SetContent(container.NewWithoutLayout(hello, newGameButton, choiceGameButton))
	sw.window.Resize(fyne.Size{Width: 400, Height: 400})
	sw.window.SetFixedSize(true)
	sw.window.ShowAndRun()
}

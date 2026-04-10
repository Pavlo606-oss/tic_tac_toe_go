package front

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type GameApp struct {
	app fyne.App
}

func NewGameApp() *GameApp {
	return &GameApp{app: app.New()}
}

package front

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/shabbyrobe/go-num"
)

type PlayingGameWindow struct {
	app    *GameApp
	window fyne.Window
}

func NewPlayingGameWindow(app *GameApp) *PlayingGameWindow {
	return &PlayingGameWindow{app: app}
}

func ChangeLinePosition(line *canvas.Line, pos1, pos2 fyne.Position) {
	line.Position1 = pos1
	line.Position2 = pos2
}

func (pgw *PlayingGameWindow) ShowNewPlayingGameWindow(idU128 num.U128) {
	pgw.window = pgw.app.app.NewWindow(fmt.Sprintf("Игра %s", idU128.String()))
	pgw.window.SetFixedSize(true)
	pgw.window.Resize(fyne.Size{Width: 400, Height: 400})
	container := container.NewWithoutLayout()
	var x, y float32
	for i := 0; i < 2; i++ {
		y += 125
		line := canvas.NewLine(color.Black)
		firstPos := fyne.NewPos(pgw.window.Canvas().Size().Width-pgw.window.Canvas().Size().Width*0.90, y)
		ChangeLinePosition(line, firstPos, fyne.NewPos(pgw.window.Canvas().Size().Width*0.90, y))
		container.Add(line)
	}
	for j := 0; j < 2; j++ {
		x += 125
		line := canvas.NewLine(color.Black)
		firstPos := fyne.NewPos(x, pgw.window.Canvas().Size().Height-pgw.window.Canvas().Size().Height*0.90)
		ChangeLinePosition(line, firstPos, fyne.NewPos(x, pgw.window.Canvas().Size().Height*0.90))
		container.Add(line)
	}
	pgw.window.SetContent(container)
	pgw.window.Show()
}

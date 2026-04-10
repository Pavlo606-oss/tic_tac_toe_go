package front

import (
	"fmt"
	"image/color"
	"tic_tac_toe/internal/logic"
	"tic_tac_toe/internal/service"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
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

func DrawBoard(container *fyne.Container, window *fyne.Window, idU128 num.U128, gs *service.GameService, width, height float32) {
	var x, y float32
	for i := 0; i < 2; i++ {
		y += 125
		line := canvas.NewLine(color.Black)
		firstPos := fyne.NewPos(width-width*0.9, y)
		ChangeLinePosition(line, firstPos, fyne.NewPos(width*0.9, y))

		container.Add(line)
	}
	for j := 0; j < 2; j++ {
		x += 125
		line := canvas.NewLine(color.Black)
		firstPos := fyne.NewPos(x, height-height*0.90)
		ChangeLinePosition(line, firstPos, fyne.NewPos(x, height*0.90))
		container.Add(line)
	}
	var buttons [3][3]*widget.Button

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			row, col := i, j
			buttons[i][j] = widget.NewButton("", func() {
				emptyInterface, _ := gs.M.Load(idU128)
				game, ok := emptyInterface.(*logic.GameLogic)
				if !ok {
					return
				}

				if game.Player == 1 {
					game.PlayerStep(uint8(row), uint8(col))
					game.ChangePlayer()
					gs.M.Store(idU128, game)
					buttons[row][col].SetText("X")
					container.Refresh()
				}
				if !(game.FullBoard() || game.CheckWinner()) {
					x, y := game.MachineStep()
					buttons[x][y].SetText("O")
					game.ChangePlayer()
					if game.FullBoard() || game.CheckWinner() {
						(*window).Close()
					}
				} else {
					(*window).Close()
				}
			})

			buttons[i][j].Move(fyne.NewPos(float32(i)*width/3, float32(j)*height/3))
			buttons[i][j].Resize(fyne.NewSize(width/3-3, height/3-3))
			container.Add(buttons[i][j])
		}
	}
}

func (pgw *PlayingGameWindow) ShowNewPlayingGameWindow(idU128 num.U128, gs *service.GameService) {
	pgw.window = pgw.app.app.NewWindow(fmt.Sprintf("Игра %s", idU128.String()))
	pgw.window.SetFixedSize(true)
	pgw.window.Resize(fyne.Size{Width: 375, Height: 375})
	container := container.NewWithoutLayout()
	DrawBoard(container, &pgw.window, idU128, gs, pgw.window.Canvas().Size().Width, pgw.window.Canvas().Size().Height)
	pgw.window.SetContent(container)
	pgw.window.Show()
	emptyInterface, _ := gs.M.Load(idU128)
	game, _ := emptyInterface.(*logic.GameLogic)
	if game.Player == -1 {
		game.MachineStep()
		game.ChangePlayer()
		gs.M.Store(idU128, game)
	}
	if game.FullBoard() || game.CheckWinner() {
		pgw.window.Close()
	}
}

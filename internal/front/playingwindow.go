package front

import (
	"image/color"
	"tic_tac_toe/internal/logic"
	"tic_tac_toe/internal/service"
	"time"

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

func WinnerName(game *logic.GameLogic) string {
	if game.CheckWinner() {
		switch -game.Player {
		case 1:
			return "Победил X!"
		case -1:
			return "Победил O!"
		}
	}
	return "Ничья!"
}

func DrawBoard(container *fyne.Container, window *fyne.Window, idU128 num.U128, gs *service.GameService, width, height float32, pgw *PlayingGameWindow) {
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
	emptyInterface, _ := gs.M.Load(idU128)
	game, _ := emptyInterface.(*logic.GameLogic)
	var buttons [3][3]*widget.Button
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			row, col := i, j
			buttons[i][j] = widget.NewButton(buttonString(game, int8(row), int8(col)), func() {
				if game.Player == 1 && !(game.FullBoard() || game.CheckWinner()) {
					if game.Board.Condition[row][col] == 0 {
						game.PlayerStep(uint8(row), uint8(col))
						buttons[row][col].SetText("X")
						gs.Db.UpdateGame(game)
					}
				}
				if !(game.FullBoard() || game.CheckWinner()) {
					if game.Player == -1 {
						x, y := game.MachineStep()
						buttons[x][y].SetText("O")
						gs.Db.UpdateGame(game)
					}
				}
				if game.FullBoard() || game.CheckWinner() {
					go func() {
						time.Sleep(2 * time.Second)
						fyne.Do(func() {
							ShowEndDialog(window, idU128, gs, pgw, game)
						})
					}()
				}
			})
			buttons[i][j].Move(fyne.NewPos(float32(i)*width/3, float32(j)*height/3))
			buttons[i][j].Resize(fyne.NewSize(width/3-3, height/3-3))
			container.Add(buttons[i][j])
		}
	}
}

func (pgw *PlayingGameWindow) ShowNewPlayingGameWindow(idU128 num.U128, gs *service.GameService) {
	SettingsPlayingWindow(pgw, idU128)
	container := container.NewWithoutLayout()
	DrawBoard(container, &pgw.window, idU128, gs, pgw.window.Canvas().Size().Width, pgw.window.Canvas().Size().Height, pgw)
	pgw.window.SetContent(container)
	pgw.window.Show()
	emptyInterface, _ := gs.M.Load(idU128)
	game, _ := emptyInterface.(*logic.GameLogic)
	if game.FullBoard() || game.CheckWinner() {
		ShowEndDialog(&pgw.window, idU128, gs, pgw, game)
	}
}

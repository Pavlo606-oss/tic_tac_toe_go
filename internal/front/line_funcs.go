package front

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func ChangeLinePosition(line *canvas.Line, pos1, pos2 fyne.Position) {
	line.Position1 = pos1
	line.Position2 = pos2
}

package front

import (
	"fmt"

	"fyne.io/fyne/v2"
	"github.com/shabbyrobe/go-num"
)

func SettingsPlayingWindow(pgw *PlayingGameWindow, idU128 num.U128) {
	pgw.window = pgw.app.app.NewWindow(fmt.Sprintf("Игра %s", idU128.String()))
	pgw.window.SetFixedSize(true)
	pgw.window.Resize(fyne.Size{Width: 375, Height: 375})
}

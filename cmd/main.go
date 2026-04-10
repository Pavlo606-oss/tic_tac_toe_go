package main

import (
	"tic_tac_toe/internal/front"
)

func main() {
	a := front.NewGameApp()
	app := front.NewStartWindow(a)
	app.ShowStartWindow()
}

package main

import (
	"github.com/fyne-io/fyne/desktop"
	"github.com/steveoc64/mapedit"
)

func main() {
	app := desktop.NewApp()
	w := app.NewWindow("MapEditor")
	w.SetContent(mapedit.NewMapedit())
	w.Show()
}

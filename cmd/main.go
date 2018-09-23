package main

import (
	"github.com/fyne-io/fyne/desktop"
	"github.com/steveoc64/mapedit"
)

func main() {
	app := desktop.NewApp()
	mapedit.NewMapedit(app, "Map Editor Demo")
}

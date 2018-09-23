package mapedit

import (
	"fmt"

	"github.com/fyne-io/fyne"
	"github.com/fyne-io/fyne/layout"
	"github.com/fyne-io/fyne/widget"
)

// NewTopbar returns a ptr to a new topbar
func (m *Mapedit) NewTopbar() fyne.CanvasObject {
	c := fyne.NewContainer()
	c.Layout = layout.NewGridLayout(12)

	c.AddObject(widget.NewButton("Open", func() {
		fmt.Println("Open File")
	}))
	c.AddObject(widget.NewButton("Save", func() {
		fmt.Println("Save File")
	}))
	c.AddObject(widget.NewButton("Save As", func() {
		fmt.Println("Save As File")
	}))
	c.AddObject(widget.NewButton("Quit", func() {
		fmt.Println("Quit")
		m.Quit()
	}))
	m.topbar = c

	return m.topbar
}

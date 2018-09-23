package mapedit

import (
	"github.com/fyne-io/fyne"
	"github.com/fyne-io/fyne/layout"
	"github.com/fyne-io/fyne/widget"
)

// Mapedit implements fyne.CanvasObject
type Mapedit struct {
	app   fyne.App
	title string
	w     fyne.Window

	// main layout
	l, topbar, body, widgetbar, widgets, options, maparea, statusbar fyne.CanvasObject
}

// Quit exits the mapeditor
func (m *Mapedit) Quit() {
	m.app.Quit()
}

// NewMapedit returns a new mapedit
func NewMapedit(app fyne.App, title string) *Mapedit {
	if title == "" {
		title = "MapEditor"
	}

	m := &Mapedit{
		app:   app,
		title: title,
		w:     app.NewWindow(title),
	}

	widgets := fyne.NewContainerWithLayout(layout.NewGridLayout(2),
		widget.NewLabel("w 1"),
		widget.NewLabel("w 2"),
		widget.NewLabel("w 3"),
		widget.NewLabel("w 4"),
		widget.NewLabel("w 5"),
		widget.NewLabel("w 6"),
	)
	options := fyne.NewContainerWithLayout(layout.NewListLayout(),
		widget.NewLabel("o 1"),
		widget.NewLabel("o 2"),
		widget.NewLabel("o 3"),
		widget.NewLabel("o 4"),
		widget.NewLabel("o 5"),
	)
	widgetbar := fyne.NewContainerWithLayout(layout.NewListLayout(),
		widgets,
		options,
	)

	maparea := fyne.NewContainer(
		widget.NewLabel("Map Area Here"),
	)
	maparea.Layout = layout.NewFixedGridLayout(fyne.NewSize(100, 100))

	body := fyne.NewContainer(
		widgetbar, // TODO - want to lock the width of this column, but not sure how :(
		maparea,
	)
	body.Layout = layout.NewFixedGridLayout(fyne.NewSize(128, 200))

	sleft := widget.NewLabel("s1")
	smid := widget.NewLabel("s2")
	sright := widget.NewLabel("s3")
	statusbar := fyne.NewContainer(sleft, smid, sright)
	statusbar.Layout = layout.NewBorderLayout(nil, nil, sleft, sright)

	l := widget.NewList(
		m.NewTopbar(),
		body,
		layout.NewSpacer(),
		statusbar,
	)
	m.w.SetContent(l)
	m.w.Show()

	return m
}

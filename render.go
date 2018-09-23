package mapedit

import (
	"github.com/fyne-io/fyne"
	"github.com/steveoc64/fyne/canvas"
)

type mapRender struct {
	bg  *canvas.Rectangle
	obj []fyne.CanvasObject
}

func newMapRender() *mapRender {
	return &mapRender{}
}

func (m *mapRender) MinSize() fyne.Size {
	return fyne.NewSize(800, 600)
}

func (m *mapRender) ApplyTheme() {}

func (m *mapRender) Layout(size fyne.Size) {
}

func (m *mapRender) Objects() []fyne.CanvasObject {
	return m.obj
}

func (m *mapRender) Refresh() {
}

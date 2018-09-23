package mapedit

import (
	"testing"

	"github.com/fyne-io/fyne/test"
	"github.com/stretchr/testify/assert"
)

func TestNewMapedit(t *testing.T) {
	app := test.NewApp()
	m := NewMapedit(app, "Test Map Editor")

	assert.NotNil(t, m)
	//assert.IsType(t, Mapedit, m, "is a Mapedit")
	//assert.IsType(t, fyne.CanvasObject, m, "is a fyne.CanvasObject")
}

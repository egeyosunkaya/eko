package netmodules

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// Module ...
type Module interface {
	GetDisplay() ui.Drawable
	GetName() string
	Run()
}

var (
	moduleList = []Module{
		NewLayer7Listener(),
	}
)

// LoadModule ...
func LoadModule(name string) Module {
	for _, mod := range moduleList {
		if mod.GetName() == name {
			return mod
		}
	}
	return nil
}

// NewLayer7Listener ...
func NewLayer7Listener() *Layer7Listener {

	box := widgets.NewParagraph()

	return &Layer7Listener{
		0,
		"Layer7Listener",
		box,
	}
}

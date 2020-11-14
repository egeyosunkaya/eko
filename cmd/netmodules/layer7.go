package netmodules

import (
	"time"

	"github.com/gizak/termui/v3/widgets"

	ui "github.com/gizak/termui/v3"
)

// Layer7Listener ...
type Layer7Listener struct {
	Port    uint32
	name    string
	Display *widgets.Paragraph
}

// Run ...
func (listener *Layer7Listener) Run() {
	go listener.startListening()
}

func (listener *Layer7Listener) startListening() {
	d, _ := time.ParseDuration("1s")
	listener.Display.Text = "AASDIHASDKJHASDJKASASDASDASasdasda"
	counter := 0
	for {

		if counter%3 == 0 {
			listener.Display.TextStyle.Fg = ui.ColorWhite
		} else if counter%3 == 1 {
			listener.Display.TextStyle.Fg = ui.ColorGreen
		} else {
			listener.Display.TextStyle.Fg = ui.ColorRed
		}
		time.Sleep(d)
		counter++
	}

}

// GetName ...
func (listener *Layer7Listener) GetName() string {
	return listener.name
}

// GetDisplay ...
func (listener *Layer7Listener) GetDisplay() ui.Drawable {
	return listener.Display
}

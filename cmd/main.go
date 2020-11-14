package main

import (
	"log"
	"time"

	"github.com/EKO/netmodules"
	ui "github.com/gizak/termui/v3"
)

func main() {

	var mod = netmodules.LoadModule("Layer7Listener")
	mod.Run()

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	mod.GetDisplay().SetRect(0, 0, 40, 40)
	ui.Render(mod.GetDisplay())

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}

	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			}
		case <-ticker:

			ui.Render(mod.GetDisplay())
		}
	}

}

package terminal

import (
	"fmt"
	"github.com/rivo/tview"
	"log"
	"time"
)

const refreshInterval = 500 * time.Millisecond

func currentTimeString() string {
	t := time.Now()
	return fmt.Sprintf(t.Format("Current time is 15:04:05"))
}

func (v View) updateTime() {
	for {
		time.Sleep(refreshInterval)
		v.app.QueueUpdateDraw(func() {
			v.view.SetText(currentTimeString())
		})
	}
}

type View struct {
	app  *tview.Application
	view *tview.Modal
}

func (v View) Start() error {
	v.app = tview.NewApplication()
	v.view = tview.NewModal().
		SetText(currentTimeString()).
		AddButtons([]string{"Quit", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Quit" {
				v.app.Stop()
			}
		})

	go v.updateTime()
	if err := v.app.SetRoot(v.view, false).Run(); err != nil {
		log.Fatalln(err)
	}

	return nil
}

package terminal

import (
	"fmt"
	"github.com/rivo/tview"
	"time"
)

func (v *View) createTaskTimerPageStart(taskName string) {
	exitCh := make(chan struct{})
	exitFunc := func() {
		exitCh <- struct{}{}
	}
	timer, _ := v.createTaskTimerPage(exitFunc)
	timer.SetTitle(taskName)
	go v.updateTime(timer, exitCh)
	v.app.SetInputCapture(nil)
	v.app.SetRoot(timer, true)
}

func (v *View) createTaskTimerPage(exitFunc func()) (*tview.Modal, error) {
	timer := tview.NewModal().
		SetText("Processing").
		AddButtons([]string{"Finish task"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Finish task" {
				//v.app.Stop()
				exitFunc()
				mainPage, _ := v.createMainPage()
				v.app.SetRoot(mainPage, true)
			}
		})
	return timer, nil
}

const refreshInterval = 500 * time.Millisecond

func (v *View) updateTime(timerBlock *tview.Modal, exit <-chan struct{}) {
	tick := time.NewTicker(refreshInterval)
	timePast := time.Now()
	for {
		select {
		case <-exit:
			return
		case <-tick.C:
			v.app.QueueUpdateDraw(func() {
				diff := time.Now().Sub(timePast)
				past := time.Time{}.Add(diff)
				str := fmt.Sprintf(past.Format(timerBlock.GetTitle() + "\nTime past - 15:04:05"))
				timerBlock.SetText(str)
			})
		}
	}
}

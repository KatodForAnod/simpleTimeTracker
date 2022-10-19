package terminal

import (
	"fmt"
	"github.com/rivo/tview"
	"log"
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
				exitFunc()
				_, err := v.controller.StopTask()
				if err != nil {
					log.Println(err)
				}
				main_Page, _ := v.createMainPage()
				v.app.SetRoot(main_Page, true)
			}
		})
	return timer, nil
}

const refreshInterval = 500 * time.Millisecond

func (v *View) updateTime(timerBlock *tview.Modal, exit <-chan struct{}) {
	tick := time.NewTicker(refreshInterval)
	timePast := time.Now()
	err := v.controller.StartTask(timerBlock.GetTitle())
	if err != nil {
		log.Println(err)
	}

	for {
		select {
		case <-exit:
			return
		case <-tick.C:
			v.app.QueueUpdateDraw(func() {
				diff := time.Now().Sub(timePast)
				past := time.Time{}.Add(diff)
				str := timerBlock.GetTitle() + fmt.Sprintf(past.Format("\nTime past - 15:04:05"))
				timerBlock.SetText(str)
			})
		}
	}
}

package terminal

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"log"
	"time"
)

type View struct {
	app *tview.Application
}

type PageName int

const (
	MenuBar     PageName = 1
	TickerBlock PageName = 2
	LastTask    PageName = 3
)

var (
	PagesHotKeys = map[PageName]tcell.Key{
		MenuBar:     tcell.KeyCtrlA,
		TickerBlock: tcell.KeyCtrlS,
		LastTask:    tcell.KeyCtrlB,
	}
	HotKeysNamed = map[tcell.Key]string{
		tcell.KeyCtrlA: "Ctrl-a",
		tcell.KeyCtrlS: "Ctrl-s",
		tcell.KeyCtrlB: "Ctrl-b",
	}
)

func (v View) Start() error {
	v.app = tview.NewApplication()
	mainPage, _ := v.createMainPage()
	if err := v.app.SetRoot(mainPage, true).Run(); err != nil {
		return err
	}
	return nil
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

func (v *View) createMainPage() (*tview.Flex, error) {
	menuBar, _ := v.createMenuBarBlock()
	taskStarter, _ := v.createTaskStarterBlock()
	lastTasks, _ := v.createLastTasksBlock()

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(menuBar, 5, 1, true).
			AddItem(taskStarter, 0, 1, false).
			AddItem(lastTasks, 0, 3, false), 0, 2, false)

	v.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case PagesHotKeys[MenuBar]:
			v.app.SetFocus(menuBar)
		case PagesHotKeys[TickerBlock]:
			v.app.SetFocus(taskStarter)
		case PagesHotKeys[LastTask]:
			v.app.SetFocus(lastTasks)
		}
		return event
	})
	return flex, nil
}

func (v *View) createMenuBarBlock() (*tview.Form, error) {
	menuBar := tview.NewForm()
	menuBar.AddButton("Главная", func() {
		log.Fatalln("u pressed main button")
	})
	menuBar.AddButton("Настройки", nil)
	menuBar.AddButton("Выход", func() {
		v.app.Stop()
	})
	menuBar.SetTitle("Top " + HotKeysNamed[PagesHotKeys[MenuBar]]).SetBorder(true)
	return menuBar, nil
}

func (v *View) createTaskStarterBlock() (*tview.Form, error) {
	taskStarter := tview.NewForm()
	inputTaskField := tview.NewInputField().SetLabel("Task:").SetFieldWidth(20)
	taskStarter.AddFormItem(inputTaskField)
	taskStarter.AddButton("Start", func() {
		exitCh := make(chan struct{})
		exitFunc := func() {
			exitCh <- struct{}{}
		}
		timer, _ := v.createTaskTimerPage(exitFunc)
		timer.SetTitle(inputTaskField.GetText())
		go v.updateTime(timer, exitCh)
		v.app.SetInputCapture(nil)
		v.app.SetRoot(timer, true)
	})
	taskStarter.SetTitle("Middle " + HotKeysNamed[PagesHotKeys[TickerBlock]]).SetBorder(true)
	taskStarter.SetHorizontal(true)
	return taskStarter, nil
}

func (v *View) createLastTasksBlock() (*tview.List, error) {
	lastTasks := tview.NewList()
	lastTasks.AddItem("Task 1", "Start: 01-01-01; End: 02-02-02; Amount 1h", 'a', nil)
	lastTasks.AddItem("Task 2", "Start: 01-01-01; End: 02-02-02; Amount 1h", 'b', nil)
	lastTasks.SetTitle("Bottom " + HotKeysNamed[PagesHotKeys[LastTask]]).SetBorder(true)
	return lastTasks, nil
}

func (v *View) createTaskTimerPage(exitFunc func()) (*tview.Modal, error) {
	timer := tview.NewModal().
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

package terminal

import (
	"github.com/rivo/tview"
)

type mainPage struct {
	createTaskTimerPage func(taskName string)
}

func (v *mainPage) createMainPage(menuBar *tview.Form, setFocus func(page PageName, primitive tview.Primitive)) (*tview.Flex, error) {
	taskStarter, _ := v.createTaskStarterBlock()
	lastTasks, _ := v.createLastTasksBlock()

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(menuBar, 5, 1, true).
			AddItem(taskStarter, 0, 1, false).
			AddItem(lastTasks, 0, 3, false), 0, 2, false)

	setFocus(MenuBar, menuBar)
	setFocus(TickerBlock, taskStarter)
	setFocus(LastTasks, lastTasks)

	return flex, nil
}

func (v *mainPage) createTaskStarterBlock() (*tview.Form, error) {
	taskStarter := tview.NewForm()
	inputTaskField := tview.NewInputField().SetLabel("Task:").SetFieldWidth(20)
	taskStarter.AddFormItem(inputTaskField)
	taskStarter.AddButton("Start", func() {
		v.createTaskTimerPage(inputTaskField.GetText())
	})
	taskStarter.SetTitle("Middle " + HotKeysNamed[PagesHotKeys[TickerBlock]]).SetBorder(true)
	taskStarter.SetHorizontal(true)
	return taskStarter, nil
}

func (v *mainPage) createLastTasksBlock() (*tview.List, error) {
	lastTasks := tview.NewList()
	lastTasks.AddItem("Task 1", "Start: 01-01-01; End: 02-02-02; Amount 1h", 'a', nil)
	lastTasks.AddItem("Task 2", "Start: 01-01-01; End: 02-02-02; Amount 1h", 'b', nil)
	lastTasks.AddItem("Task 3", "Start: 01-01-01; End: 02-02-02; Amount 1h", 'b', nil)
	lastTasks.AddItem("Task 4", "Start: 01-01-01; End: 02-02-02; Amount 1h", 'b', nil)
	lastTasks.AddItem("Task 5", "Start: 01-01-01; End: 02-02-02; Amount 1h", 'b', nil)
	lastTasks.AddItem("More", "Load more ↓", 'b', nil)
	lastTasks.SetTitle("Bottom " + HotKeysNamed[PagesHotKeys[LastTasks]]).SetBorder(true)
	return lastTasks, nil
}
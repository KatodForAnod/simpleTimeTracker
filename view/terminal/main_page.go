package terminal

import (
	"fmt"
	"github.com/rivo/tview"
	"log"
	"simpleTimeTracker/pkg/controller"
	"simpleTimeTracker/pkg/models"
	"time"
)

type mainPage struct {
	controller          controller.App
	createTaskTimerPage func(taskName string)
}

func (v *mainPage) createMainPage(menuBar *tview.Form, setFocus func(page PageName, primitive tview.Primitive)) (*tview.Flex, error) {
	taskStarter, _ := v.createTaskStarterBlock()
	lastTasks, _ := v.createLastTasksBlock()

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(menuBar, 5, 1, false).
			AddItem(taskStarter, 5, 1, true).
			AddItem(lastTasks, 0, 3, false), 0, 2, true)

	setFocus(MenuBar, menuBar)
	setFocus(TickerBlock, taskStarter)
	setFocus(LastTasks, lastTasks)

	return flex, nil
}

func (v *mainPage) createTaskStarterBlock() (*tview.Form, error) {
	taskStarter := tview.NewForm()
	inputTaskField := tview.NewInputField().SetLabel("Task:").SetFieldWidth(35)
	taskStarter.AddFormItem(inputTaskField)
	taskStarter.AddButton("Start", func() {
		v.createTaskTimerPage(inputTaskField.GetText())
	})
	taskStarter.SetTitle("Middle " + HotKeysNamed[PagesHotKeys[TickerBlock]]).SetBorder(true)
	taskStarter.SetHorizontal(true)
	return taskStarter, nil
}

func (v *mainPage) createLastTasksBlock() (*tview.List, error) {
	params := models.ReqTaskParams{
		Start: time.Now().Add(-(time.Hour * 24 * 7)), // one week
		Name:  "",
		Limit: 10,
	}

	tasks, err := v.controller.SearchTasks(params)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	lastTasks := tview.NewList()
	lastTasks.SetSelectedFocusOnly(true)

	for _, task := range tasks {
		start := task.Start.Format("2006-02-01")
		end := task.End.Format("2006-02-01")
		amount := task.End.Sub(task.Start)
		secondaryText := fmt.Sprintf("Start: %s; End: %s; Amount %s", start, end, amount.String())
		lastTasks.AddItem(task.Name, secondaryText, 'a', nil)
	}

	lastTasks.SetTitle("Bottom " + HotKeysNamed[PagesHotKeys[LastTasks]]).SetBorder(true)
	return lastTasks, nil
}

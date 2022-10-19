package terminal

import (
	"fmt"
	"github.com/rivo/tview"
	"simpleTimeTracker/pkg/models"
)

type searchPage struct {
	searchBlock *tview.Form
	tasksBlock  *tview.List
	amountBlock *tview.TextView

	tasks    []models.Task
	currPage int
}

func (p *searchPage) createSearchPage(menuBar *tview.Form, setFocus func(page PageName, primitive tview.Primitive)) (*tview.Flex, error) {
	p.tasks = []models.Task{{Name: "task0"},
		{Name: "task1"}, {Name: "task2"},
		{Name: "task3"}, {Name: "task4"},
		{Name: "task5"}, {Name: "task6"}}

	p.tasksBlock = tview.NewList()
	p.tasksBlock.SetTitle("Result " + HotKeysNamed[PagesHotKeys[SearchBlockResults]]).SetBorder(true)
	p.initSearchBlock()

	p.amountBlock = tview.NewTextView()
	p.amountBlock.SetBorderPadding(0, 0, 1, 1)
	p.amountBlock.SetText("From - To: 2022-01-01 - 2022-02-02; Amount: 10min")
	p.amountBlock.SetBorder(true)

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(p.searchBlock, 4, 1, true).
			AddItem(p.tasksBlock, 0, 3, false).
			AddItem(p.amountBlock, 3, 3, false), 0, 2, false)

	setFocus(SearchBlockParams, p.searchBlock)
	setFocus(SearchBlockResults, p.tasksBlock)
	return flex, nil
}

func (p *searchPage) initSearchBlock() error {
	p.searchBlock = tview.NewForm()
	taskName := tview.NewInputField().SetLabel("Task:").SetFieldWidth(20)
	p.searchBlock.AddFormItem(taskName)

	dropdown := tview.NewDropDown().
		SetLabel("Duration: ").
		SetOptions([]string{"Day", "Week", "Month"}, func(text string, index int) {
			switch text {
			case "Day":
				p.updateTasks()
			case "Week":
				p.updateTasks()
			case "Month":
				p.updateTasks()
			default:
			}
		})
	p.searchBlock.AddFormItem(dropdown)
	p.searchBlock.SetBorder(true)
	p.searchBlock.SetTitle("Search " + HotKeysNamed[PagesHotKeys[SearchBlockParams]])
	p.searchBlock.SetItemPadding(0)
	p.searchBlock.SetBorderPadding(0, 0, 0, 0)
	return nil
}

func (p *searchPage) updateTasks() {
	if len(p.tasks) <= p.currPage*5 {
		return
	}

	p.tasksBlock.Clear()
	if p.currPage != 0 {
		p.tasksBlock.AddItem("Back", "Return back ↑", 'b', func() {
			p.currPage = p.currPage - 1
			p.updateTasks()
		})
	}
	for i := p.currPage * 5; i < (p.currPage+1)*5 && i < len(p.tasks); i++ {
		start := p.tasks[i].Start.Format("2006-02-01")
		end := p.tasks[i].End.Format("2006-02-01")
		amount := "x" // temporary
		secondaryText := fmt.Sprintf("Start: %s; End: %s, Amount: %s", start, end, amount)
		p.tasksBlock.AddItem(p.tasks[i].Name, secondaryText, 'a', nil)
	}
	if (p.currPage+1)*5 < len(p.tasks) {
		p.tasksBlock.AddItem("More", "Load more ↓", 'b', func() {
			p.currPage = p.currPage + 1
			p.updateTasks()
		})
	}
}

package terminal

import (
	"fmt"
	"github.com/rivo/tview"
	"log"
	"simpleTimeTracker/pkg/models"
	"time"
)

type searchPage struct {
	searchBlock *tview.Form
	tasksBlock  *tview.List
	amountBlock *tview.TextView

	tasks       []models.Task
	searchTasks func(params models.ReqTaskParams) ([]models.Task, error)

	currPage int
}

func (p *searchPage) createSearchPage(menuBar *tview.Form,
	setFocus func(page PageName, primitive tview.Primitive),
	searchTasks func(params models.ReqTaskParams) ([]models.Task, error)) (*tview.Flex, error) {
	p.searchTasks = searchTasks

	p.tasksBlock = tview.NewList()
	p.tasksBlock.SetSelectedFocusOnly(true)
	p.tasksBlock.SetTitle("Result " + HotKeysNamed[PagesHotKeys[SearchBlockResults]]).SetBorder(true)
	_ = p.initSearchBlock()

	p.amountBlock = tview.NewTextView()
	p.amountBlock.SetBorderPadding(0, 0, 1, 1)
	p.amountBlock.SetText("From - To: 2022-01-01 - 2022-02-02; Amount: 10min")
	p.amountBlock.SetBorder(true)

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(menuBar, 5, 1, false).
			AddItem(p.searchBlock, 4, 1, true).
			AddItem(p.tasksBlock, 0, 3, false).
			AddItem(p.amountBlock, 3, 3, false), 0, 2, true)

	setFocus(SearchBlockParams, p.searchBlock)
	setFocus(SearchBlockResults, p.tasksBlock)
	setFocus(MenuBar, menuBar)
	return flex, nil
}

func (p *searchPage) initSearchBlock() error {
	p.searchBlock = tview.NewForm()
	taskName := tview.NewInputField().SetLabel("Task:").SetFieldWidth(20)
	p.searchBlock.AddFormItem(taskName)

	dropdown := tview.NewDropDown().
		SetLabel("Duration: ").
		SetOptions([]string{"Day", "Week", "Month"}, func(text string, index int) {
			params := models.ReqTaskParams{
				Name:  taskName.GetText(),
				Limit: 1000,
			}
			const day = time.Hour * 24
			const week = day * 7
			const month = day * 31
			switch text {
			case "Day":
				params.Start = time.Now().Add(-day)
			case "Week":
				params.Start = time.Now().Add(-week)
			case "Month":
				params.Start = time.Now().Add(-month)
			default:
				return
			}
			tasks, err := p.searchTasks(params)
			if err != nil {
				log.Println(err)
				return
			}
			p.tasks = tasks
			p.initTasks()
			_ = p.updateAmountBlock(params.Start, time.Now())
		})
	p.searchBlock.AddFormItem(dropdown)
	p.searchBlock.SetBorder(true)
	p.searchBlock.SetTitle("Search " + HotKeysNamed[PagesHotKeys[SearchBlockParams]])
	p.searchBlock.SetItemPadding(0)
	p.searchBlock.SetBorderPadding(0, 0, 0, 0)
	return nil
}

func (p *searchPage) updateAmountBlock(start, end time.Time) error {
	var amount time.Duration
	for _, task := range p.tasks {
		amount += task.End.Sub(task.Start)
	}

	p.amountBlock.SetText(fmt.Sprintf(`From - To: %s - %s; Amount: %s`,
		start.Format("2006-02-01"), end.Format("2006-02-01"), amount))
	return nil
}

func (p *searchPage) initTasks() {
	p.currPage = 0
	p.updateTasks()
}

func (p *searchPage) updateTasks() {
	const countOfTasksView = 8
	p.tasksBlock.Clear()
	if p.currPage != 0 {
		p.tasksBlock.AddItem("Back", "Return back ↑", 'b', func() {
			p.currPage = p.currPage - 1
			p.updateTasks()
		})
	}
	for i := p.currPage * countOfTasksView; i < (p.currPage+1)*countOfTasksView && i < len(p.tasks); i++ {
		start := p.tasks[i].Start.Format("2006-02-01")
		end := p.tasks[i].End.Format("2006-02-01")
		amount := "x" // temporary
		secondaryText := fmt.Sprintf("Start: %s; End: %s, Amount: %s", start, end, amount)
		p.tasksBlock.AddItem(p.tasks[i].Name, secondaryText, 'a', nil)
	}
	if (p.currPage+1)*countOfTasksView < len(p.tasks) {
		p.tasksBlock.AddItem("More", "Load more ↓", 'b', func() {
			p.currPage = p.currPage + 1
			p.updateTasks()
		})
	}
}

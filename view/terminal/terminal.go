package terminal

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"log"
	"simpleTimeTracker/pkg/controller"
)

type View struct {
	controller controller.App

	app     *tview.Application
	menuBar *tview.Form
}

func (v *View) Init(app controller.App) {
	v.controller = app
}

func (v *View) Start() error {
	v.app = tview.NewApplication()
	main_page, _ := v.createMainPage()
	if err := v.app.SetRoot(main_page, true).Run(); err != nil {
		return err
	}
	return nil
}

func (v *View) ShutDown() error {
	v.controller.StopTask() //add check is task running
	v.app.Stop()
	return nil
}

func (v *View) createMainPage() (*tview.Flex, error) {
	menuBar, _ := v.createMenuBarBlock()
	mainPageStruct := mainPage{createTaskTimerPage: v.createTaskTimerPageStart, controller: v.controller} //controller singleton?
	mainPageObj, _ := mainPageStruct.createMainPage(menuBar, v.createNewFuncInputCapture())
	return mainPageObj, nil
}

func (v *View) createNewFuncInputCapture() func(page PageName, primitive tview.Primitive) {
	v.app.SetInputCapture(nil) // ??can be a problem
	focusFunc := func(page PageName, primitive tview.Primitive) {
		f := v.app.GetInputCapture()
		v.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			switch event.Key() {
			case PagesHotKeys[page]:
				v.app.SetFocus(primitive)
			}
			if f != nil {
				f(event)
			}
			return event
		})
	}

	return focusFunc
}

func (v *View) createSearchPage() (*tview.Flex, error) {
	searchPageObj := searchPage{}
	menuBar, _ := v.createMenuBarBlock()
	page, _ := searchPageObj.createSearchPage(menuBar, v.createNewFuncInputCapture())
	return page, nil
}

func (v *View) createMenuBarBlock() (*tview.Form, error) {
	if v.menuBar != nil {
		return v.menuBar, nil
	}
	menuBar := tview.NewForm()
	v.menuBar = menuBar

	menuBar.AddButton("Главная", func() {
		main_page, _ := v.createMainPage() //TODO memory use fix
		v.app.SetRoot(main_page, true)
	})
	menuBar.AddButton("Поиск", func() {
		search_page, _ := v.createSearchPage() //TODO memory use fix
		v.app.SetRoot(search_page, true)
	})
	menuBar.AddButton("Настройки", nil)
	menuBar.AddButton("Выход", func() {
		err := v.controller.ShutDown()
		if err != nil {
			log.Println(err)
			return
		}
		v.app.Stop()
	})
	menuBar.SetTitle("Top " + HotKeysNamed[PagesHotKeys[MenuBar]]).SetBorder(true)
	return menuBar, nil
}

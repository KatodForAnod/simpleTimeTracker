package terminal

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"log"
)

type View struct {
	app *tview.Application
}

func (v View) Start() error {
	v.app = tview.NewApplication()
	main_page, _ := v.createMainPage()
	if err := v.app.SetRoot(main_page, true).Run(); err != nil {
		return err
	}
	return nil
}

func (v *View) createMainPage() (*tview.Flex, error) {
	menuBar, _ := v.createMenuBarBlock() //fix memory use
	mainPageStruct := mainPage{createTaskTimerPage: v.createTaskTimerPageStart}
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
	menuBar, _ := v.createMenuBarBlock() //fix memory use
	page, _ := searchPageObj.createSearchPage(menuBar, v.createNewFuncInputCapture())
	return page, nil
}

func (v *View) createMenuBarBlock() (*tview.Form, error) {
	menuBar := tview.NewForm()
	menuBar.AddButton("Главная", func() {
		main_page, _ := v.createMainPage()
		if err := v.app.SetRoot(main_page, true).Run(); err != nil {
			log.Println(err)
		}
	})
	menuBar.AddButton("Настройки", nil)
	menuBar.AddButton("Выход", func() {
		v.app.Stop()
	})
	menuBar.SetTitle("Top " + HotKeysNamed[PagesHotKeys[MenuBar]]).SetBorder(true)
	return menuBar, nil
}

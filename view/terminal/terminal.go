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
	menuBar, _ := v.createMenuBarBlock()
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
	mainPageStruct := mainPage{createTaskTimerPage: v.createTaskTimerPageStart}
	mainPageObj, _ := mainPageStruct.createMainPage(menuBar, focusFunc)
	return mainPageObj, nil
}

/*func (v View) Start() error {
	v.app = tview.NewApplication()

	s := searchPage{}
	flex, _ := s.createSearchPage(v.app)

	if err := v.app.SetRoot(flex, true).Run(); err != nil {
		return err
	}
	return nil
}*/

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

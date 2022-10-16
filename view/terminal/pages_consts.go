package terminal

import "github.com/gdamore/tcell/v2"

type PageName int

const (
	MenuBar            PageName = 1
	TickerBlock        PageName = 2
	LastTasks          PageName = 3
	SearchBlockParams  PageName = 4
	SearchBlockResults PageName = 5
)

var (
	PagesHotKeys = map[PageName]tcell.Key{
		MenuBar:            tcell.KeyCtrlA,
		TickerBlock:        tcell.KeyCtrlS,
		LastTasks:          tcell.KeyCtrlB,
		SearchBlockParams:  tcell.KeyCtrlH,
		SearchBlockResults: tcell.KeyCtrlF,
	}
	HotKeysNamed = map[tcell.Key]string{
		tcell.KeyCtrlA: "Ctrl-a",
		tcell.KeyCtrlS: "Ctrl-s",
		tcell.KeyCtrlB: "Ctrl-b",
		tcell.KeyCtrlH: "Ctrl-h",
		tcell.KeyCtrlF: "Ctrl-f",
	}
)

package terminal

import "github.com/gdamore/tcell/v2"

type PageName int

const (
	MenuBar          PageName = 1
	TickerBlock      PageName = 2
	LastTask         PageName = 3
	SearchBlock      PageName = 4
	SearchBlockTasks PageName = 5
)

var (
	PagesHotKeys = map[PageName]tcell.Key{
		MenuBar:          tcell.KeyCtrlA,
		TickerBlock:      tcell.KeyCtrlS,
		LastTask:         tcell.KeyCtrlB,
		SearchBlock:      tcell.KeyCtrlH,
		SearchBlockTasks: tcell.KeyCtrlF,
	}
	HotKeysNamed = map[tcell.Key]string{
		tcell.KeyCtrlA: "Ctrl-a",
		tcell.KeyCtrlS: "Ctrl-s",
		tcell.KeyCtrlB: "Ctrl-b",
		tcell.KeyCtrlH: "Ctrl-h",
		tcell.KeyCtrlF: "Ctrl-f",
	}
)

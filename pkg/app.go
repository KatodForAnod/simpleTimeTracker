package pkg

import (
	"fmt"
	"simpleTimeTracker/pkg/controller"
)

type typeDB string

const (
	SQLiteDB   typeDB = "sqlite"
	SysFilesDB typeDB = "sysFiles"
)

func InitApp(db typeDB) (controller.App, error) {
	switch db {
	case SQLiteDB:
	case SysFilesDB:
	default:
		return nil, fmt.Errorf("InitApp err: db %s does not support", db)
	}

	panic("realize me!!!")
}

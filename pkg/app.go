package pkg

import (
	"fmt"
	"log"
	"simpleTimeTracker/pkg/controller"
	"simpleTimeTracker/pkg/db/sqlite"
)

type typeDB string

const (
	SQLiteDB   typeDB = "sqlite"
	SysFilesDB typeDB = "sysFiles"
)

func InitApp(db typeDB) (controller.App, error) {
	var app controller.App

	switch db {
	case SQLiteDB:
		dbSQL := sqlite.SqlLite{}
		if err := dbSQL.InitDataBase(); err != nil {
			log.Println(err)
			return nil, err
		}
		if err := dbSQL.CreateTables(); err != nil {
			log.Println(err)
			return nil, err
		}
		contr := controller.InitController(&dbSQL)
		app = &contr
	case SysFilesDB:
	default:
		return nil, fmt.Errorf("InitApp err: db %s does not support", db)
	}

	return app, nil
}

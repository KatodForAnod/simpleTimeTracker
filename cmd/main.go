package main

import (
	"fmt"
	"log"
	"simpleTimeTracker/pkg/controller"
	"simpleTimeTracker/pkg/db/sqlite"
	"simpleTimeTracker/view/terminal"
)

func main() {
	log.SetFlags(log.Lshortfile)
	lite := sqlite.SqlLite{}

	err := lite.InitDataBase()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = lite.CreateTables()
	if err != nil {
		fmt.Println(err)
		return
	}

	controller := controller.InitController(&lite)
	view := terminal.View{}
	view.Init(&controller)
	view.Start()
}

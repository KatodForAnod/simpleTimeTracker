package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"simpleTimeTracker/pkg/controller"
	"simpleTimeTracker/pkg/db/sqlite"
	"simpleTimeTracker/view/terminal"
	"syscall"
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

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-c
		view.ShutDown()
	}()
	view.Start()
}

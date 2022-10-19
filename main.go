package main

import (
	"log"
	"os"
	"os/signal"
	"simpleTimeTracker/pkg"
	"simpleTimeTracker/view/terminal"
	"syscall"
)

func main() {
	log.SetFlags(log.Lshortfile)

	app, err := pkg.InitApp(pkg.SQLiteDB)
	if err != nil {
		log.Println(err)
		return
	}

	view := terminal.View{}
	view.Init(app)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGHUP)
	go func() {
		<-c
		view.ShutDown()
	}()
	view.Start()
}

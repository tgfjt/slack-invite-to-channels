package main

import (
	"github.com/tgfjt/slack-invite-to-channels/lib"
	"github.com/zserge/lorca"
)

func main() {
	// TODO: is server goes down... ui
	// Start goroutine
	go lib.Server()

	var ui lorca.UI
	ui, _ = lorca.New("", "", 1280, 800)
	defer ui.Close()

	ui.Load("http://localhost:12345")

	<-ui.Done()
}

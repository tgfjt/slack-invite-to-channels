package main

import (
	"github.com/zserge/lorca"
)

func main() {
	server()

	var ui lorca.UI
	ui, _ = lorca.New("", "", 1024, 800)

	ui.Load("http://localhost:12345")

	<-ui.Done()
	defer ui.Close()

}

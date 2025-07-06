package main

import (
	"gopad/helpers"
	"gopad/ui"
	"os"
)

func main() {

	flagHandler := helpers.FlagHandlerFunc()

	flagHandler.HandleFlags()

	args := os.Args[1:] // Exclude the first argument which is the program name

	if len(args) > 0 {
		ui.NewWindow(args[0])
	} else {
		ui.NewWindow("")
	}
}

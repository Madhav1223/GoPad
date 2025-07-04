package main

import (
	"gopad/internal/ui"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) > 0 && args[0] == "version" {
		println("GoPad Version 2.0")
		return
	}
	if len(args) > 0 && args[0] == "help" {
		println("Usage: gopad [version|help][`file_path`]")
		println("  version: Show the version of GoPad")
		println("  help: Show this help message")
		println("  file_path: Open a specific file in GoPad")
		// take the first argument as a file path if provided
		return
	}
	ui.NewWindow(args)
}

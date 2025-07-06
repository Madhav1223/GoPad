package helpers

import (
	"flag"
	"fmt"
	"os"
)

// FlagHandler struct holds the flags and provides methods to handle them.
type FlagHandler struct {
	HelpFlag    *bool
	VersionFlag *bool
}

// newFlagHandler initializes the FlagHandler struct with the required flags (private).
func newFlagHandler() *FlagHandler {
	helpFlag := flag.Bool("help", false, "Show help message")
	versionFlag := flag.Bool("version", false, "Show version")
	flag.Parse()

	return &FlagHandler{
		HelpFlag:    helpFlag,
		VersionFlag: versionFlag,
	}
}

// handleFlags parses the flags and performs actions based on the user's input (private).
func (f *FlagHandler) handleFlags() {
	if *f.VersionFlag {
		fmt.Println("GoPad Version 2.0")
		os.Exit(0)
		return
	}

	if *f.HelpFlag {
		f.printHelp()
		os.Exit(0)
		return
	}
}

// printHelp displays the help message for the program (private).
func (f *FlagHandler) printHelp() {
	fmt.Println("Usage: gopad [--version] [--help] [file_path]")
	fmt.Println("  --version: Show the version of GoPad")
	fmt.Println("  --help: Show this help message")
	fmt.Println("  file_path: Open a specific file in GoPad")
}

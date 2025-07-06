package helpers

func FlagHandlerFunc() *FlagHandler {
	return newFlagHandler()
}

func (f *FlagHandler) HandleFlags() {
	f.handleFlags()
}

func (f *FlagHandler) PrintHelp() {
	f.printHelp()
}

func OpenFile(path string) (string, error) {
	return openFile(path)
}

func CreateNewFile(path string) error {
	return createNewFile(path)
}

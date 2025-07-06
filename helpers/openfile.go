package helpers

import (
	"os"
)

// openFile reads the content of the file and returns it as a string (private).
func openFile(path string) (string, error) {
	//  check if file exists or if path is blank or not
	if path == "" {
		return "", nil // Return empty string if no path is provided
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// If the file does not exist, create a new file with some initial content
		if err := createNewFile(path); err != nil {
			return "", err
		}
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// createNewFile creates a new file and writes some initial content to it (private).
func createNewFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("New file created!\n")
	if err != nil {
		return err
	}
	return nil
}

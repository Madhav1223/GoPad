# GoPad 📝

**GoPad** is a modern, lightweight, and feature-rich text editor built with [Go](https://golang.org/) and the [Fyne](https://fyne.io/) GUI toolkit. Designed for developers, students, and anyone who needs a simple yet powerful text editor, GoPad combines a clean user interface with essential editing tools. Advanced features are planned for future releases.

GoPad supports multiple text editing functionalities, including regex-based search, syntax highlighting for Go, file management, and more. Whether you're writing code or notes, GoPad aims to provide a productive, seamless experience.

---

## ✨ Features

### Core Functionality

- [x] **New File**: Start a fresh text file.
- [x] **Open File**: Open an existing file.
- [x] **Save and Save As**: Save files with different names or locations.
- [ ] **Unsaved Changes Warning**: Prevent accidental loss of data by warning users of unsaved changes.
- [ ] **Cut, Undo, Redo**: Basic editing functions for better control.
- [ ] **Word Wrap Toggle**: Enable/disable word wrapping for easier reading and editing.

### User Experience Enhancements

- [ ] **Search & Replace** (with **Regex** support):
  - **Regex Search**: Advanced searching using regular expressions for complex pattern matching.
  - **Find All Matches**: Highlight all matches of the search term in the document.
  - **Replace with Regex**: Replace text with regex-based replacement patterns (including capture groups).
- [ ] **Find in Files**: Search across multiple files in a project.
- [ ] **Multi-tab Support**: Open and edit multiple files simultaneously in different tabs.
- [ ] **Auto-save**: Set automatic saving intervals to avoid data loss.
- [ ] **Status Bar**: Display cursor position, file status, and word count.
- [ ] **Line Numbers**: Display line numbers in the editor for better navigation.

### Customization & Personalization

- [ ] **Font Customization**: Change font type and size.
- [ ] **Keybinding Customization**: Define custom keyboard shortcuts for various actions.
- [ ] **Theme Preferences**: Save and load Dark/Light mode settings per session.
- [ ] **Split View**: Edit two files side by side in a single window.

### Extras

- [ ] **Markdown Preview**: Preview markdown files within the editor.
- [ ] **Code Snippets**: Define and insert reusable code snippets.
- [ ] **Bracket Matching & Highlighting**: Highlight matching braces, parentheses, and brackets in code.

---

## 🚀 Quick Start

### Prerequisites

To run **GoPad**, you need:

- **Go 1.20** or higher
- [Fyne](https://fyne.io/) GUI toolkit

### Installation

1. **Install Fyne:**

   ```bash
   go install fyne.io/fyne/v2/cmd/fyne@latest
   go get fyne.io/fyne/v2
   ```

2. **Clone the Repository:**

   ```bash
   git clone https://github.com/madhav1223/GoNote.git
   cd GoNote
   ```

3. **Run the Application:**

   ```bash
   go run main.go
   ```

---

## 🗂️ Project Structure

```
GoPad/
├── main.go         # Entry point for the application
├── src/            # (Optional) Additional layout and GUI-related code
├── .env            # Theme and environment settings
├── go.mod          # Go module definition
└── README.md       # Project documentation
```

---

## 🛠️ Usage

- **New File:** File → New or `Ctrl+N`
- **Open File:** File → Open or `Ctrl+O`
- **Save File:** File → Save or `Ctrl+S`
- **Save As:** File → Save As
- **Copy/Paste:** Edit → Copy/Paste or `Ctrl+C`/`Ctrl+V`
- **Search & Replace:** `Ctrl+F` for search, `Ctrl+H` for advanced search & replace (with regex)
- **Toggle Themes:** Switch between Dark and Light mode in the status bar

---

## 🙌 Contributing

We welcome contributions from the community! Whether it's bug fixes, new features, or improvements to documentation, feel free to open a pull request.

For major changes, please open an issue first to discuss what you’d like to change. We encourage feedback and suggestions to make GoPad a better tool for everyone!

---

## 📚 Resources

- [Fyne Documentation](https://docs.fyne.io/)
- [Go Documentation](https://golang.org/doc/)

---

## 📝 License

GoPad is an open-source project licensed under the MIT License. It is intended for educational purposes and personal use.

---

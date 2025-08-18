package main

import (
	"fmt"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
)

// === Start of data model ===

type DirectoryItemType string

const (
	File   DirectoryItemType = "file"
	Folder DirectoryItemType = "folder"
)

type DirectoryItem struct {
	contents []DirectoryItem
	name     string
	itemType DirectoryItemType
}

func NewFolder(name string, contents []DirectoryItem) DirectoryItem {
	return DirectoryItem{
		itemType: Folder,
		name:     name,
		contents: contents,
	}
}

func NewFile(name string) DirectoryItem {
	return DirectoryItem{
		itemType: File,
		name:     name,
	}
}

// === End of data model ===

// === Start of View Model ===

type TraversableDirectory struct {
	DirectoryItem
	selectedElement  string
	currentDirectory string
}

func (dirModel *TraversableDirectory) Init() tea.Cmd {
	return nil
}

func (dirModel *TraversableDirectory) View() string {
	view := fmt.Sprintf("Current Directory: %s\n\n", dirModel.currentDirectory)

	// only print current directory for now
	for _, dirItem := range dirModel.contents {
		if dirItem.itemType == File {
			view += fmt.Sprintf("%s\n", dirItem.name)
		} else {
			view += fmt.Sprintf("%s/\n", dirItem.name)
		}
	}

	return view
}

func (dirModel *TraversableDirectory) Update(
	msg tea.Msg,
) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return dirModel, tea.Quit
		}
	}

	return dirModel, nil
}

func NewTraversableDirectory(cwd string) *TraversableDirectory {
	rawDirectoryContents, err := os.ReadDir(cwd)
	if err != nil {
		fmt.Println("Error: Could not get directory contents", err)
		os.Exit(1)
	}

	// convert directory contents into DirectoryItems
	directoryContents := []DirectoryItem{}
	for _, content := range rawDirectoryContents {
		if content.IsDir() {
			directoryContents = append(
				directoryContents,
				NewFolder(content.Name(), []DirectoryItem{}),
			)
		} else {
			directoryContents = append(
				directoryContents,
				NewFile(content.Name()),
			)
		}
	}

	directoryItem := NewFolder(filepath.Base(cwd), directoryContents)

	return &TraversableDirectory{
		DirectoryItem:    directoryItem,
		selectedElement:  "test",
		currentDirectory: cwd,
	}
}

// === End of View Model ===

// === Start of main ===

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error: Could not get current working directory", err)
		os.Exit(1)
	}

	p := tea.NewProgram(NewTraversableDirectory(cwd))

	if err := p.Start(); err != nil {
		fmt.Println("Error: Could not start program", err)
		os.Exit(1)
	}
}

// === End of main ===

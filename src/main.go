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
	name     string
	itemType DirectoryItemType
}

func NewFolder(name string, contents []DirectoryItem) *DirectoryItem {
	return &DirectoryItem{
		itemType: Folder,
		name:     name,
	}
}

func NewFile(name string) *DirectoryItem {
	return &DirectoryItem{
		itemType: File,
		name:     name,
	}
}

// === End of data model ===

// === Start of View Model ===

type TraversableDirectory struct {
	contents []DirectoryItem
	// Lists the present working directory (pwd).
	pwd string
	// Denotes selected item in current directory.
	itemIndex int
}

func (dirModel *TraversableDirectory) Init() tea.Cmd {
	return nil
}

func (dirModel *TraversableDirectory) View() string {
	view := fmt.Sprintf("Current Directory: %s\n\n", dirModel.pwd)

	// only print current directory for now
	for index, dirItem := range dirModel.contents {
		item := ""
		if dirItem.itemType == File {
			item = fmt.Sprintf(" %s", dirItem.name)
		} else {
			item = fmt.Sprintf("\033[35m %s/\033[0m", dirItem.name)
		}

		cursor := " "
		underlineStart := ""
		underlineEnd := ""
		if index == dirModel.itemIndex {
			cursor = ">"
			underlineStart = "\033[1;4m"
			underlineEnd = "\033[0m"
		}
		view += fmt.Sprintf(
			"%d %s %s %s %s\n",
			index+1,
			cursor,
			underlineStart,
			item,
			underlineEnd,
		)
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
		case "j", "down":
			if dirModel.itemIndex < len(dirModel.contents)-1 {
				dirModel.itemIndex++
			}
		case "k", "up":
			if dirModel.itemIndex > 0 {
				dirModel.itemIndex--
			}
		case "h", "left":
			parentDir := filepath.Dir(dirModel.pwd)
			entries, err := os.ReadDir(parentDir)
			if err != nil {
				fmt.Println("Error: Could not read from new directory", err)
				return dirModel, tea.Quit
			}

			dirModel.pwd = parentDir
			dirModel.itemIndex = 0
			dirModel.contents = convertDirEntriesToDirectoryItems(
				entries,
			)
		case "l", "right":
			if len(dirModel.contents) == 0 {
				break
			}

			currentItem := dirModel.contents[dirModel.itemIndex]
			if currentItem.itemType != Folder {
				break
			}

			newPath := filepath.Join(dirModel.pwd, currentItem.name)
			entries, err := os.ReadDir(newPath)
			if err != nil {
				fmt.Println("Error: Could not read from new directory", err)
				return dirModel, tea.Quit
			}

			dirModel.pwd = newPath
			dirModel.contents = convertDirEntriesToDirectoryItems(
				entries,
			)
			dirModel.itemIndex = 0
		case "o":
			if len(dirModel.contents) == 0 {
				break
			}

			lastDirFile := os.Getenv("FT_LAST_DIR")
			if lastDirFile != "" {
				os.WriteFile(
					lastDirFile,
					[]byte(fmt.Sprintf("cd %s\n", dirModel.pwd)),
					0755,
				)
			}
			return dirModel, tea.Quit
		}
	}

	return dirModel, nil
}

func convertDirEntriesToDirectoryItems(
	rawDirectoryContents []os.DirEntry,
) []DirectoryItem {
	directoryContents := []DirectoryItem{}
	for _, content := range rawDirectoryContents {
		if content.IsDir() {
			directoryContents = append(
				directoryContents,
				*NewFolder(content.Name(), []DirectoryItem{}),
			)
		} else {
			directoryContents = append(
				directoryContents,
				*NewFile(content.Name()),
			)
		}
	}
	return directoryContents
}

func NewTraversableDirectory(cwd string) *TraversableDirectory {
	rawDirectoryContents, err := os.ReadDir(cwd)
	if err != nil {
		fmt.Println("Error: Could not get directory contents", err)
		os.Exit(1)
	}

	directoryContents := convertDirEntriesToDirectoryItems(
		rawDirectoryContents,
	)

	return &TraversableDirectory{
		contents:  directoryContents,
		pwd:       cwd,
		itemIndex: 0,
	}
}

// === End of View Model ===

// === Start of main ===

func main() {
	cwd, err := os.Getwd()
	fmt.Println(cwd)
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

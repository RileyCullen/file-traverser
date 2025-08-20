package traversabledirectory

import (
	dm "file-traverser/src/directory-model"
	"fmt"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
)

// Update handles key inputs and updates the model accordingly.
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
			if currentItem.ItemType != dm.Folder {
				break
			}

			newPath := filepath.Join(dirModel.pwd, currentItem.Name)
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

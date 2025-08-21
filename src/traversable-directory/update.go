package traversabledirectory

import (
	dm "file-traverser/src/directory-model"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// Update handles key inputs and updates the model accordingly.
func (dirModel *ViewModel) Update(
	msg tea.Msg,
) (tea.Model, tea.Cmd) {
	actionType := dirModel.textBufferAction.actionType
	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()
		switch key {
		case "q":
			return dirModel, tea.Quit
		case "j", "down":
			if dirModel.itemIndex < len(dirModel.contents)-1 {
				dirModel.itemIndex = updateItemIndex(
					dirModel,
					false,
				)
			}
			dirModel.textBufferAction = *newNoopTextBufferAction()
		case "k", "up":
			if dirModel.itemIndex > 0 {
				dirModel.itemIndex = updateItemIndex(
					dirModel,
					true,
				)
			}
			dirModel.textBufferAction = *newNoopTextBufferAction()
		case "h", "left":
			parentDir := filepath.Dir(dirModel.pwd)
			entries, err := os.ReadDir(parentDir)
			if err != nil {
				fmt.Println("Error: Could not read from new directory", err)
				return dirModel, tea.Quit
			}

			prevPwd := dirModel.pwd
			dirModel.pwd = parentDir
			dirModel.contents = convertDirEntriesToDirectoryItems(
				entries,
			)
			dirModel.itemIndex = getDirectoryItemIndex(
				filepath.Base(prevPwd),
				dirModel,
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
		case ":":
			switch actionType {
			case noop:
				dirModel.textBufferAction = *newExplicitLineChangeTextBufferAction()
			default:
				dirModel.textBufferAction = *newNoopTextBufferAction()
			}
		case "enter":
			switch actionType {
			case explicitLineChange:
				dirModel.itemIndex = updateItemIndex(
					dirModel,
					false,
				)
			case noop:
				return dirModel, nil
			}
			dirModel.textBufferAction = *newNoopTextBufferAction()
		default:
			if key >= "0" && key <= "9" {
				switch actionType {
				case noop:
					dirModel.textBufferAction = *newRelativeLineChangeTextBufferAction(
						[]string{key},
					)
				case relativeLineChange:
				case explicitLineChange:
					dirModel.textBufferAction.buffer = append(
						dirModel.textBufferAction.buffer,
						key,
					)
				default:
					fmt.Println("Error: Invalid buffer action type")
					return dirModel, tea.Quit
				}
			}
		}
	}

	return dirModel, nil
}

func updateItemIndex(
	dirModel *ViewModel,
	isMovingUp bool,
) int {
	index := getIndexFromBufferAction(dirModel, isMovingUp)
	maxIndex := len(dirModel.contents) - 1

	if index > maxIndex {
		return maxIndex
	}

	if index < 0 {
		return 0
	}

	return index
}

func getIndexFromBufferAction(
	dirModel *ViewModel,
	isMovingUp bool,
) int {
	actionType := dirModel.textBufferAction.actionType
	currentIndex := dirModel.itemIndex

	amount := 1
	if actionType == relativeLineChange || actionType == explicitLineChange {
		var err error
		amount, err = parseLineChangeBuffer(dirModel.textBufferAction.buffer)
		if err != nil {
			return currentIndex
		}

		if actionType == explicitLineChange {
			return amount - 1
		}
	}

	if isMovingUp {
		return currentIndex - amount
	}
	return currentIndex + amount
}

func parseLineChangeBuffer(buffer []string) (int, error) {
	lineChangeAmount := strings.Join(buffer, "")
	return strconv.Atoi(lineChangeAmount)
}

func getDirectoryItemIndex(name string, dirModel *ViewModel) int {
	for i, dirItem := range dirModel.contents {
		if dirItem.Name == name {
			return i
		}
	}
	return 0
}

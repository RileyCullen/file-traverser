package traversabledirectory

import (
	dm "file-traverser/src/directory-model"
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// View of terminal output given current dirModel state.
func (dirModel *ViewModel) View() string {
	view := displayPresentWorkingDirectory(dirModel)
	view += displayDirectoryContents(dirModel)

	if dirModel.textBufferAction.actionType != noop {
		view += "\n" + lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			Width(50).
			Render(strings.Join(dirModel.textBufferAction.buffer, ""))
	}

	return view
}

func displayPresentWorkingDirectory(dirModel *ViewModel) string {
	return fmt.Sprintf("Current Directory: %s\n\n", dirModel.pwd)
}

func displayDirectoryContents(dirModel *ViewModel) string {
	directoryContents := ""
	maxPadding := len(strconv.Itoa(len(dirModel.contents)))

	for index, dirItem := range dirModel.contents {
		item := getDisplayValueForCurrentItem(dirItem)

		cursor, formatStart, formatEnd := getFormatForCurrentLine(
			index,
			dirModel,
		)

		directoryContents += fmt.Sprintf(
			"%s %s %s %s %s\n",
			getLineNumber(index+1, maxPadding),
			cursor,
			formatStart,
			item,
			formatEnd,
		)
	}
	return directoryContents
}

type color string

const (
	purple color = "\033[35m"
)

const reset = "\033[0m"

func getDisplayValueForCurrentItem(dirItem dm.DirectoryItem) string {
	item := ""
	if dirItem.ItemType == dm.File {
		item = fmt.Sprintf(" %s", dirItem.Name)
	} else {
		item = fmt.Sprintf("%s %s/%s", purple, dirItem.Name, reset)
	}
	return item
}

const boldUnderline = "\033[1;4m"

func getFormatForCurrentLine(
	index int,
	dirModel *ViewModel,
) (string, string, string) {
	cursor := " "
	formatStart := ""
	formatEnd := ""
	if index == dirModel.itemIndex {
		cursor = ">"
		formatStart = boldUnderline
		formatEnd = reset
	}
	return cursor, formatStart, formatEnd
}

func getLineNumber(number int, maxPadding int) string {
	lineNumber := fmt.Sprintf("%d", number)
	actualPadding := maxPadding - len(strconv.Itoa(number))
	for range actualPadding {
		lineNumber += " "
	}
	return lineNumber
}

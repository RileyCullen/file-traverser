package traversabledirectory

import (
	dm "file-traverser/src/directory-model"
	"fmt"
)

// View of terminal output given current dirModel state.
func (dirModel *TraversableDirectory) View() string {
	view := displayPresentWorkingDirectory(dirModel)
	view += displayDirectoryContents(dirModel)
	return view
}

func displayPresentWorkingDirectory(dirModel *TraversableDirectory) string {
	return fmt.Sprintf("Current Directory: %s\n\n", dirModel.pwd)
}

func displayDirectoryContents(dirModel *TraversableDirectory) string {
	directoryContents := ""
	for index, dirItem := range dirModel.contents {
		item := getDisplayValueForCurrentItem(dirItem)

		cursor, formatStart, formatEnd := getFormatForCurrentLine(
			index,
			dirModel,
		)

		directoryContents += fmt.Sprintf(
			"%d %s %s %s %s\n",
			index+1,
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
	dirModel *TraversableDirectory,
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

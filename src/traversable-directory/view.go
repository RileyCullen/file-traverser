package traversabledirectory

import (
	dm "file-traverser/src/directory-model"
	"fmt"
)

func (dirModel *TraversableDirectory) View() string {
	view := fmt.Sprintf("Current Directory: %s\n\n", dirModel.pwd)

	// only print current directory for now
	for index, dirItem := range dirModel.contents {
		item := ""
		if dirItem.ItemType == dm.File {
			item = fmt.Sprintf(" %s", dirItem.Name)
		} else {
			item = fmt.Sprintf("\033[35m %s/\033[0m", dirItem.Name)
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

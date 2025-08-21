package traversabledirectory

import (
	dm "file-traverser/src/directory-model"
	"fmt"
	"os"
)

// ViewModel is a model for viewing a directory.
type ViewModel struct {
	// Directory contents.
	contents []dm.DirectoryItem
	// Lists the present working directory (pwd).
	pwd string
	// Denotes selected item in current directory.
	itemIndex int
	// Since tea only supports single character presses, to support :<num> and
	// 10<j | k> operations, we need to store historical key strokes in memory.
	textBufferAction textBufferAction
}

// NewViewModel initializes view model based on current directory.
func NewViewModel(cwd string) *ViewModel {
	rawDirectoryContents, err := os.ReadDir(cwd)
	if err != nil {
		fmt.Println("Error: Could not get directory contents", err)
		os.Exit(1)
	}

	directoryContents := convertDirEntriesToDirectoryItems(
		rawDirectoryContents,
	)

	return &ViewModel{
		contents:         directoryContents,
		pwd:              cwd,
		itemIndex:        0,
		textBufferAction: *newNoopTextBufferAction(),
	}
}

func convertDirEntriesToDirectoryItems(
	rawDirectoryContents []os.DirEntry,
) []dm.DirectoryItem {
	directoryContents := []dm.DirectoryItem{}
	for _, content := range rawDirectoryContents {
		if content.IsDir() {
			directoryContents = append(
				directoryContents,
				*dm.NewFolder(content.Name(), []dm.DirectoryItem{}),
			)
		} else {
			directoryContents = append(
				directoryContents,
				*dm.NewFile(content.Name()),
			)
		}
	}
	return directoryContents
}

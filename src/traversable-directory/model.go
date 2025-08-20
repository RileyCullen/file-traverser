package traversabledirectory

import (
	dm "file-traverser/src/directory-model"
	"fmt"
	"os"
)

// TraversableDirectory is a model for viewing a directory.
type TraversableDirectory struct {
	// Directory contents.
	contents []dm.DirectoryItem
	// Lists the present working directory (pwd).
	pwd string
	// Denotes selected item in current directory.
	itemIndex int
}

// NewTraversableDirectory initializes view model based on current directory.
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

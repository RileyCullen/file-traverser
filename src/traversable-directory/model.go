package traversabledirectory

import (
	dm "file-traverser/src/directory-model"
	"fmt"
	"os"
)

type TraversableDirectory struct {
	contents []dm.DirectoryItem
	// Lists the present working directory (pwd).
	pwd string
	// Denotes selected item in current directory.
	itemIndex int
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

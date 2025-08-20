package directorymodel

// DirectoryItemType is an enum denoting type of an item in the directory.
type DirectoryItemType string

const (
	File   DirectoryItemType = "file"
	Folder DirectoryItemType = "folder"
)

// DirectoryItem models an item in a directory, which can be a file or folder.
type DirectoryItem struct {
	Name     string
	ItemType DirectoryItemType
}

// NewFolder is a constructor function that creates a new DirectoryItem
// of type Folder.
func NewFolder(name string, contents []DirectoryItem) *DirectoryItem {
	return &DirectoryItem{
		ItemType: Folder,
		Name:     name,
	}
}

// NewFile is a constructor function that creates a new DirectoryItem of
// type file.
func NewFile(name string) *DirectoryItem {
	return &DirectoryItem{
		ItemType: File,
		Name:     name,
	}
}

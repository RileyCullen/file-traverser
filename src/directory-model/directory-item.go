package directorymodel

type DirectoryItemType string

const (
	File   DirectoryItemType = "file"
	Folder DirectoryItemType = "folder"
)

type DirectoryItem struct {
	Name     string
	ItemType DirectoryItemType
}

func NewFolder(name string, contents []DirectoryItem) *DirectoryItem {
	return &DirectoryItem{
		ItemType: Folder,
		Name:     name,
	}
}

func NewFile(name string) *DirectoryItem {
	return &DirectoryItem{
		ItemType: File,
		Name:     name,
	}
}

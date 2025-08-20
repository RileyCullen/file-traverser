package main

import (
	td "file-traverser/src/traversable-directory"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	cwd, err := os.Getwd()
	fmt.Println(cwd)
	if err != nil {
		fmt.Println("Error: Could not get current working directory", err)
		os.Exit(1)
	}

	p := tea.NewProgram(td.NewViewModel(cwd))

	if err := p.Start(); err != nil {
		fmt.Println("Error: Could not start program", err)
		os.Exit(1)
	}
}

package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"cal"
	"log"
)

func main() {
	file, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	defer file.Close()

	if _, err := tea.NewProgram(cal.CreateParentModel()).Run(); err != nil {
		log.Fatal(err)
	}
}

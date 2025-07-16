package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"cal"
)

func main() {
	// file, err := tea.LogToFile("debug.log", "debug")
	// if err != nil {
	// 	log.Fatalf("err: %v", err)
	// }
	// defer file.Close()

	my_model := cal.ParentModel{
		Is_in_day_view: false,
		Dayview: cal.DayViewModel{},
		Monthview: *cal.CreateMonthViewModel(),

	}
	if _, err := tea.NewProgram(my_model).Run(); err != nil {
		// log.Fatal(err)
	}
}

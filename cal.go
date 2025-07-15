package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	lg "github.com/charmbracelet/lipgloss"
)

// lipgloss styles 
// -------------------------------------------

var dayCellNormalModelStyle = lg.NewStyle().
		BorderStyle(lg.NormalBorder()).
		Width(12).
		Height(4)

var dayCellFocusedModelStyle = lg.NewStyle().
		Inherit(dayCellNormalModelStyle).
		BorderForeground(lipgloss.Color("69"))

// ---------------------------------------------

// bubbletea shit
// ---------------------------------------------

type model struct {
	daycellcount uint
	focused_day uint
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
				case "ctrl+c":
					return m, tea.Quit
				case "l":
					if m.focused_day < m.daycellcount - 1 {
						m.focused_day++
					}
				case "h":
					if m.focused_day > 0 {
						m.focused_day--
					}
		}
	}
	return m, nil
}

func (m model) View() string {
	days := make([]string, m.daycellcount)
	for index := range(m.daycellcount) {
		if index == m.focused_day {
			days[index] = dayCellFocusedModelStyle.Render()
		} else {
			days[index] = dayCellNormalModelStyle.Render()
		}
	}
	return lg.JoinHorizontal(lg.Left, days...)
}

// ---------------------------------------------

func main() {
	program := tea.NewProgram(model{daycellcount: 7})
	if _, err := program.Run(); err != nil {
		// idk how to handle errors lmao
	}
}

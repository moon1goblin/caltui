package cal

import (
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
)

type DayViewModel struct {
	parent_model_ptr *ParentModel
}

var dayviewstyle = lg.NewStyle().
		Border(lg.RoundedBorder()).
		Padding()

func (m *DayViewModel) Init() tea.Cmd {
	return nil
}

func (m *DayViewModel) View() string {
	content := "this is the day view :_"
	return dayviewstyle.Render(content)
}

func (m *DayViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+[":
			m.parent_model_ptr.Is_in_day_view = false
		}
	}
	return m, nil
}

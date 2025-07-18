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
	content := lg.JoinVertical(
		lg.Left,
		m.parent_model_ptr.focused_day_time.Format("Mon Jan 2 2006"),
		lg.NewStyle().Padding(5).Render("this is the day view :_)"),
	)
	return dayviewstyle.Render(content)
}

// ughhhhh fuck this shit
func (m *DayViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "h":
			m.parent_model_ptr.ChangeFocusedDay(-1)
		case "l":
			m.parent_model_ptr.ChangeFocusedDay(1)
		case "esc", "ctrl+[":
			m.parent_model_ptr.is_in_day_view = false
		}
	}
	return m, nil
}

package cal

import tea "github.com/charmbracelet/bubbletea"

type DayViewModel struct {
}

// func (m dayViewModel) Init() tea.Cmd {
// 	return nil
// }

func (m DayViewModel) ViewDayView() string {
	return "this is the day view :_)"
}

func (m DayViewModel) UpdateDayView(parentmodel *ParentModel, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
				case "ctrl+c", "q":
				case "esc", "ctrl+[":
					parentmodel.Is_in_day_view = false
			}
	}
	return parentmodel, nil
}

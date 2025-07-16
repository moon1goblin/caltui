package cal

import (
	tea "github.com/charmbracelet/bubbletea"
)

const calwidth_g uint = 7
const calheight_g uint = 4

type ParentModel struct {
	Is_in_day_view bool
	Dayview DayViewModel
	Monthview MonthViewModel
}

// // [begintime, endtime]
// func loadDays(begintime time.Time, endtime time.Time) tea.Cmd {
// 	return func() tea.Msg {
// 		// ask the server for that if theres a connection
// 		// if not ask sqlite
// 		// i am not doing that shit rn though
// 	}
// }

func (m ParentModel) Init() tea.Cmd {
	return nil
}

func (m ParentModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.Is_in_day_view {
		return m.Dayview.UpdateDayView(&m, msg)
	} else {
		return m.Monthview.UpdateMonthView(&m, msg)
	}
}

func (m ParentModel) View() string {
	if m.Is_in_day_view {
		return m.Dayview.ViewDayView()
	} else {
		return m.Monthview.ViewMonthView()
	}
}

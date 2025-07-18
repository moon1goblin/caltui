package cal

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	overlay "github.com/rmhubbert/bubbletea-overlay"
)

// cal dimensions
const dimension_calwidth_g uint = 7
const dimension_calheight_g uint = 4

// TODO: change dim with terminal size

type ParentModel struct {
	// TODO: make this an enum
	is_in_day_view bool
	dayview *DayViewModel
	monthview *MonthViewModel
	overlay tea.Model

	cur_time time.Time
	// monday 0 sunday 6
	cur_weekday uint
	focused_day_time time.Time
}

// // [begintime, endtime]
// func loadDays(begintime time.Time, endtime time.Time) tea.Cmd {
// 	return func() tea.Msg {
// 		// ask the server for that if theres a connection
// 		// if not ask sqlite
// 		// i am not doing that shit rn though
// 	}
// }

func CreateParentModel() *ParentModel {
	m := ParentModel{}
	m.is_in_day_view = false

	m.cur_time = time.Now()
	m.focused_day_time = m.cur_time
	// times weekdays start from sunday, eww
	m.cur_weekday = uint((int(m.cur_time.Weekday()) + 6) % 7)

	// WTF: & or no &
	m.dayview = &DayViewModel{&m}
	m.monthview = CreateMonthViewModel(&m)

	m.overlay = overlay.New(
		m.dayview,
		m.monthview,
		overlay.Center,
		overlay.Center,
		0,
		0,
	)
	return &m
}

func (m *ParentModel) Init() tea.Cmd {
	return nil
}

func (m *ParentModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// to make sure we can always quit, maybe reconsider this later
	switch msg := msg.(type){
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd

	switch m.is_in_day_view {
	case true:
		// ok fuck this
		// ignoring tea.Model return because dayview.Update() is interfaced by pointer
		_, cmd = m.dayview.Update(msg)
	case false:
		_, cmd = m.monthview.Update(msg)
		// m.monthview = mod
	}

	return m, cmd
}

func (m *ParentModel) View() string {
	if m.is_in_day_view {
		return m.overlay.View()
	} else {
		return m.monthview.View()
	}
}

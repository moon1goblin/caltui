package cal

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	overlay "github.com/rmhubbert/bubbletea-overlay"
)

const calwidth_g uint = 7
const calheight_g uint = 4

type ParentModel struct {
	// TODO: make this an enum
	Is_in_day_view bool
	dayview tea.Model
	monthview tea.Model
	overlay tea.Model
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
	m.Is_in_day_view = true

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

	switch m.Is_in_day_view {
	case true:
		m.dayview, cmd = m.dayview.Update(msg)
	case false:
		m.monthview, cmd = m.monthview.Update(msg)
	}
	log.Println(m.Is_in_day_view)

	return m, cmd
}

func (m *ParentModel) View() string {
	if m.Is_in_day_view {
		return m.overlay.View()
	} else {
		return m.monthview.View()
	}
}

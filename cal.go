package main

import (
	"strconv"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
)

// -------------------------------------------

const calwidth uint = 7
const calheight uint = 5

type daycell struct {
	date uint
	month string
	year uint
	// also data about events here
}

// -------------------------------------------

var dayCellNormalModelStyle = lg.NewStyle().
		BorderStyle(lg.NormalBorder()).
		Width(12).
		Height(4)

var dayCellFocusedModelStyle = lg.NewStyle().
		Inherit(dayCellNormalModelStyle).
		BorderForeground(lg.Color("69"))

var dayCellTodayModelStyle = lg.NewStyle().
		Inherit(dayCellNormalModelStyle).
		BorderForeground(lg.Color("9"))

// ---------------------------------------------

type model struct {
	cur_time time.Time
	cur_weekday uint // monday 0 sunday 6
	focused_day_x uint
	focused_day_y uint
	upper_left_day time.Time
	loaded_daycells []daycell
}

// // [begintime, endtime]
// func loadDays(begintime time.Time, endtime time.Time) tea.Cmd {
// 	return func() tea.Msg {
// 		// ask the server for that if theres a connection
// 		// if not ask sqlite
// 		// i am not doing that shit rn though
// 	}
// }

func createModel() *model {
	m := model{}
	m.cur_time = time.Now()
	m.cur_weekday = uint((int(m.cur_time.Weekday()) + 6) % 7)
	m.upper_left_day = m.cur_time.AddDate(0, 0, - (7 + int(m.cur_weekday)))
	m.focused_day_x = m.cur_weekday
	m.focused_day_y = 1

	// // load days 2 month before and after, if scroll far enough ask for more
	// for i := range calheight * calwidth {
	// }
	return &m
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
					if m.focused_day_x < calwidth - 1 {
						m.focused_day_x++
					}
				case "h":
					if m.focused_day_x > 0 {
						m.focused_day_x--
					}
				case "j":
					if m.focused_day_y == calheight - 1 {
						m.upper_left_day = m.upper_left_day.AddDate(0, 0, 7)
					} else if m.focused_day_y < calheight - 1 {
						m.focused_day_y++
					}
				case "k":
					if m.focused_day_y == 0 {
						m.upper_left_day = m.upper_left_day.AddDate(0, 0, -7)
					} else if m.focused_day_y > 0 {
						m.focused_day_y--
					}
		}
	}
	return m, nil
}

func (m model) View() string {
	month := make([]string, calheight)
	week := make([]string, calwidth)

	cur_time := m.upper_left_day
	for cury := range calheight {
		for curx := range calwidth {
			if m.focused_day_x == curx && m.focused_day_y == cury {
				week[curx] = dayCellFocusedModelStyle.Render(strconv.Itoa(cur_time.Day()))
					// strconv.Itoa(cur_time.Day())
			} else if cur_time.Equal(m.cur_time) {
				week[curx] = dayCellTodayModelStyle.Render(strconv.Itoa(cur_time.Day()))
			} else {
				week[curx] = dayCellNormalModelStyle.Render(strconv.Itoa(cur_time.Day()))
			}
			cur_time = cur_time.AddDate(0, 0, 1)
		}
		month[cury] = lg.JoinHorizontal(lg.Left, week...)
	}
	return lg.JoinVertical(lg.Left, month...)
}

// ---------------------------------------------

func main() {
	// file, err := tea.LogToFile("debug.log", "debug")
	// if err != nil {
	// 	log.Fatalf("err: %v", err)
	// }
	// defer file.Close()

	if _, err := tea.NewProgram(*createModel()).Run(); err != nil {
		// log.Fatal(err)
	}
}

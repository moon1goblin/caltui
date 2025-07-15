package main

import (
	"strconv"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
)

// -------------------------------------------

const calwidth_g uint = 7
const calheight_g uint = 5

// couldnt make this constant but whatever
var weekdays_rendered = func() string {
	var weekdays_g = []string{
		"Mon",
		"Tue",
		"Wed",
		"Thu",
		"Fri",
		"Sat",
		"Sun",
	}
	for index, weekday := range(weekdays_g) {
		// add some padding, want 12 wide, 12-3 / 2 = 4 and 5
		weekdays_g[index] = "     " + weekday + "      "
	}
	return lg.JoinHorizontal(lg.Left, weekdays_g...)
}

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
		BorderStyle(lg.ThickBorder()).
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
	focused_day_time time.Time

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
	m.focused_day_time = m.cur_time

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
				case "ctrl+c", "q":
					return m, tea.Quit
				case "h":
					if m.focused_day_x > 0 {
						m.focused_day_x--
						m.focused_day_time = m.focused_day_time.AddDate(0, 0, -1)
					}
				case "l":
					if m.focused_day_x < calwidth_g - 1 {
						m.focused_day_x++
						m.focused_day_time = m.focused_day_time.AddDate(0, 0, 1)
					}
				case "k":
					if m.focused_day_y == 0 {
						m.upper_left_day = m.upper_left_day.AddDate(0, 0, -7)
					} else if m.focused_day_y > 0 {
						m.focused_day_y--
					}
					m.focused_day_time = m.focused_day_time.AddDate(0, 0, -7)
				case "j":
					if m.focused_day_y == calheight_g - 1 {
						m.upper_left_day = m.upper_left_day.AddDate(0, 0, 7)
					} else if m.focused_day_y < calheight_g - 1 {
						m.focused_day_y++
					}
					m.focused_day_time = m.focused_day_time.AddDate(0, 0, 7)
		}
	}
	return m, nil
}

func (m model) View() string {
	date_at_the_top_rendered := 
		"                                          " +
		m.focused_day_time.Format("Mon Jan 2 2006")
	alltogethernow := []string{date_at_the_top_rendered, "", weekdays_rendered()}

	month := make([]string, calheight_g)
	week := make([]string, calwidth_g)

	cur_time := m.upper_left_day
	for cury := range calheight_g {
		for curx := range calwidth_g {
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

	alltogethernow = append(alltogethernow, month...)
	return lg.JoinVertical(lg.Left, alltogethernow...)
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

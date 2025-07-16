package cal

import (
	"strconv"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
)

var (
	dayCellNormalModelStyle = lg.NewStyle().
		BorderStyle(lg.NormalBorder()).
		Width(14).
		Height(5)

	dayCellFocusedModelStyle = lg.NewStyle().
		Inherit(dayCellNormalModelStyle).
		BorderStyle(lg.ThickBorder()).
		BorderForeground(lg.Color("69"))
	dayCellTodayModelStyle = lg.NewStyle().
		Inherit(dayCellNormalModelStyle).
		BorderForeground(lg.Color("9"))
)

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

// type daycell struct {
// 	date uint
// 	month string
// 	year uint
// 	// also data about events here
// }

type MonthViewModel struct {
	// determines current window
	is_in_day_view bool

	cur_time time.Time
	// monday 0 sunday 6
	cur_weekday uint

	focused_day_x uint
	focused_day_y uint
	focused_day_time time.Time

	upper_left_day time.Time

	// this isnt necesary rn but
	// loaded_daycells []daycell
}

func CreateMonthViewModel() *MonthViewModel {
	m := MonthViewModel{}

	m.cur_time = time.Now()
	m.focused_day_time = m.cur_time
	// times weekdays start from sunday, eww
	m.cur_weekday = uint((int(m.cur_time.Weekday()) + 6) % 7)

	m.upper_left_day = m.cur_time.AddDate(0, 0, - (7 + int(m.cur_weekday)))
	m.focused_day_x = m.cur_weekday
	m.focused_day_y = 1

	// later: load days 2 month before and after, if scroll far enough ask for more

	return &m
}

func (m MonthViewModel) UpdateMonthView(parentmodel *ParentModel, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
				case "ctrl+c", "q":
					return parentmodel, tea.Quit
				case "h":
					if parentmodel.Monthview.focused_day_x > 0 {
						parentmodel.Monthview.focused_day_x--
						parentmodel.Monthview.focused_day_time = parentmodel.Monthview.focused_day_time.AddDate(0, 0, -1)
					}
				case "l":
					if parentmodel.Monthview.focused_day_x < calwidth_g - 1 {
						parentmodel.Monthview.focused_day_x++
						parentmodel.Monthview.focused_day_time = parentmodel.Monthview.focused_day_time.AddDate(0, 0, 1)
					}
				case "k":
					if parentmodel.Monthview.focused_day_y == 0 {
						parentmodel.Monthview.upper_left_day = parentmodel.Monthview.upper_left_day.AddDate(0, 0, -7)
					} else if parentmodel.Monthview.focused_day_y > 0 {
						parentmodel.Monthview.focused_day_y--
					}
					parentmodel.Monthview.focused_day_time = parentmodel.Monthview.focused_day_time.AddDate(0, 0, -7)
				case "j":
					if parentmodel.Monthview.focused_day_y == calheight_g - 1 {
						parentmodel.Monthview.upper_left_day = parentmodel.Monthview.upper_left_day.AddDate(0, 0, 7)
					} else if parentmodel.Monthview.focused_day_y < calheight_g - 1 {
						parentmodel.Monthview.focused_day_y++
					}
					parentmodel.Monthview.focused_day_time = parentmodel.Monthview.focused_day_time.AddDate(0, 0, 7)
				case "enter", "i":
					parentmodel.Is_in_day_view = true
		}
	}
	return parentmodel, nil
}

func (m MonthViewModel) ViewMonthView() string {
	date_at_the_top_rendered := 
		"                                          " +
		m.focused_day_time.Format("Mon Jan 2 2006")
	alltogethernow := []string{date_at_the_top_rendered, "", weekdays_rendered()}

	// month is made up of calheight_g weeks but shhh dont tell anyone
	month := make([]string, calheight_g)
	week := make([]string, calwidth_g)

	cur_time := m.upper_left_day
	for cury := range calheight_g {
		for curx := range calwidth_g {
			if m.focused_day_x == curx && m.focused_day_y == cury {
				week[curx] = dayCellFocusedModelStyle.Render(
					strconv.Itoa(cur_time.Day()))
			} else if cur_time.Equal(m.cur_time) {
				week[curx] = dayCellTodayModelStyle.Render(
					strconv.Itoa(cur_time.Day()))
			} else {
				week[curx] = dayCellNormalModelStyle.Render(
					strconv.Itoa(cur_time.Day()))
			}
			cur_time = cur_time.AddDate(0, 0, 1)
		}
		month[cury] = lg.JoinHorizontal(lg.Left, week...)
	}

	alltogethernow = append(alltogethernow, month...)
	return lg.JoinVertical(lg.Left, alltogethernow...)
}

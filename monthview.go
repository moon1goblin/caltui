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
	
	weekDayStyle = lg.NewStyle().
			PaddingLeft(6).
			PaddingRight(7)
	
	dateAtTheTopStyle = lg.NewStyle().
			PaddingLeft(48)
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
		weekdays_g[index] = weekDayStyle.Render(weekday)
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
	// FIXME: is there a curcular dependancy here? idk
	parent_model_ptr *ParentModel
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

func CreateMonthViewModel(parent_model_ptr *ParentModel) *MonthViewModel {
	cur_time := time.Now()
	// times weekdays start from sunday, eww
	cur_weekday := uint((int(cur_time.Weekday()) + 6) % 7)

	m := MonthViewModel{
		parent_model_ptr: parent_model_ptr,
		cur_time: cur_time,
		focused_day_time: cur_time,
		cur_weekday: cur_weekday,
		upper_left_day: cur_time.AddDate(0, 0, - (7 + int(cur_weekday))),
		focused_day_x: cur_weekday,
		focused_day_y: 1,
	}
	// later: load days 2 month before and after, if scroll far enough ask for more

	return &m
}

func (m *MonthViewModel) Init() tea.Cmd {
	return nil
}

func (m *MonthViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "h":
			if m.focused_day_x  > 0 {
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
		case "enter", "i":
			m.parent_model_ptr.Is_in_day_view = true
		}
	}
	return m, nil
}

func (m *MonthViewModel) View() string {
	alltogethernow := []string{
		dateAtTheTopStyle.Render(m.focused_day_time.Format("Mon Jan 2 2006")),
		"",
		weekdays_rendered(),
	}

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

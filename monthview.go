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
	parent_model_ptr *ParentModel

	// from 0 to 6
	focused_day_x uint
	// from 0 to dimension_height or whatever i called that
	focused_day_y uint
	upper_left_day time.Time

	// this isnt necesary rn but
	// loaded_daycells []daycell
}

func CreateMonthViewModel(parent_model_ptr *ParentModel) *MonthViewModel {

	m := MonthViewModel{
		parent_model_ptr: parent_model_ptr,
		upper_left_day: parent_model_ptr.cur_time.AddDate(0, 0, - (7 + int(parent_model_ptr.cur_weekday))),
		focused_day_x: parent_model_ptr.cur_weekday,
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
			if m.focused_day_x > 0 {
				m.parent_model_ptr.ChangeFocusedDay(-1)
			}
		case "l":
			if m.focused_day_x < dimension_calwidth_g - 1 {
				m.parent_model_ptr.ChangeFocusedDay(1)
			}
		case "k":
			m.parent_model_ptr.ChangeFocusedDay(-7)
		case "j":
			m.parent_model_ptr.ChangeFocusedDay(7)
		case "enter", "i":
			m.parent_model_ptr.is_in_day_view = true
		}
	}
	return m, nil
}

func (m *MonthViewModel) View() string {
	alltogethernow := []string{
		dateAtTheTopStyle.Render(m.parent_model_ptr.focused_day_time.Format("Mon Jan 2 2006")),
		"",
		weekdays_rendered(),
	}

	// month is made up of calheight_g weeks but shhh dont tell anyone
	month := make([]string, dimension_calheight_g)
	week := make([]string, dimension_calwidth_g)

	cur_time := m.upper_left_day
	for cury := range dimension_calheight_g {
		for curx := range dimension_calwidth_g {
			if m.focused_day_x == curx && m.focused_day_y == cury {
				week[curx] = dayCellFocusedModelStyle.Render(
					strconv.Itoa(cur_time.Day()))
			} else if cur_time.Equal(m.parent_model_ptr.cur_time) {
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

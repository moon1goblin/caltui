package cal

func (m *ParentModel) CalcAndSetFocusedDayXY() {
	// idk lol im tired
	// so if focused day is upper left days_from_upper_left = 0
	days_from_upper_left := uint(m.focused_day_time.Sub(m.monthview.upper_left_day).Hours() / 24)

	// log.Print("focused day: ", m.focused_day_time.Day(),
	// 	" upper left day: ", m.monthview.upper_left_day.Day(),
	// 	" days from upper left day: ", days_from_upper_left,
	// )

	m.monthview.focused_day_x = days_from_upper_left % 7
	m.monthview.focused_day_y = days_from_upper_left / 7
}

func (m *ParentModel) ChangeFocusedDay(increment int) {
	m.focused_day_time = m.focused_day_time.AddDate(0, 0, increment)

	// scrolling
	if increment < 0 && m.monthview.focused_day_y == 0 {
		m.monthview.upper_left_day = m.monthview.upper_left_day.AddDate(0, 0, -7)
	} else if increment > 0 && m.monthview.focused_day_y == dimension_calheight_g - 1 {
		m.monthview.upper_left_day = m.monthview.upper_left_day.AddDate(0, 0, 7)
	}

	// xy on monthview
	m.CalcAndSetFocusedDayXY()
}

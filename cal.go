package main

import (
	"log"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
)

const calwidth uint = 7
const calheight uint = 5

type daycell struct {
	date uint
	month string
	year uint
	// also data about events here
}

// lipgloss styles 
// -------------------------------------------

var dayCellNormalModelStyle = lg.NewStyle().
		BorderStyle(lg.NormalBorder()).
		Width(12).
		Height(4)

var dayCellFocusedModelStyle = lg.NewStyle().
		Inherit(dayCellNormalModelStyle).
		BorderForeground(lg.Color("69"))

// ---------------------------------------------

// bubbletea shit
// ---------------------------------------------

type model struct {
	focused_day_x uint
	focused_day_y uint
	daycells [calheight*calwidth]daycell
}

func createModel() *model {
	m := model{}
	for index := range len(m.daycells) {
		m.daycells[index].date = uint(index)
	}
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
					if m.focused_day_y < calheight - 1 {
						m.focused_day_y++
					}
				case "k":
					if m.focused_day_y > 0 {
						m.focused_day_y--
					}
		}
	}
	return m, nil
}

func (m model) View() string {
	month := make([]string, calheight)
	week := make([]string, calwidth)

	var indtotal uint = 0
	for cury := range calheight {
		for curx := range calwidth {
			if curdate := m.daycells[indtotal].date; 
				cury == m.focused_day_y && curx == m.focused_day_x {
				week[curx] = dayCellFocusedModelStyle.Render(strconv.Itoa(int(curdate)))
			} else {
				week[curx] = dayCellNormalModelStyle.Render(strconv.Itoa(int(curdate)))
			}
			indtotal++
			log.Println("heyyyyy babe")
		}
		month[cury] = lg.JoinHorizontal(lg.Left, week...)
	}
	return lg.JoinVertical(lg.Left, month...)
}

// ---------------------------------------------

func main() {
	file, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	defer file.Close()

	if _, err := tea.NewProgram(*createModel()).Run(); err != nil {
		log.Fatal(err)
	}
}

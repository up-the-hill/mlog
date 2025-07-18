package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type State int

const BASE_PATH = "/tmp/"

const (
	StateStart State = iota
	StateBook
	StateMovie
	StateQuote
	StateMusing
)

var stateNames = map[int]string{
	1: "book",
	2: "movie",
	3: "quote",
	4: "musing",
}

type book struct {
	Name string
	Date time.Time
}

type movie struct {
	Name string
	Date time.Time
}

type quote struct {
	Quote string
	Date  time.Time
}

type musing struct {
	Musing string
	Date   time.Time
}

func appendEntry(filename string, newEntry any) {
	f, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	entryJson, _ := json.Marshal(newEntry)
	f.Write(append(entryJson, '\n'))
}

type model struct {
	cursor    int
	choices   []string
	state     int
	textinput textinput.Model
}

func initialModel() model {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	return model{
		choices:   []string{"Book", "Movie", "Quote", "Thoughts/Musings"},
		state:     int(StateStart),
		textinput: ti,
	}
}

func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("mlog")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch m.state {
	case int(StateStart):
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down", "j":
				if m.cursor < len(m.choices)-1 {
					m.cursor++
				}
			case "enter", " ":
				m.state = m.cursor + 1
			}
		}
	default:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "ctrl+c":
				return m, tea.Quit
			case "enter":
				// select file to write to
				f := BASE_PATH + stateNames[m.state] + "s.ndjson"
				// format output
				switch m.state {
				case int(StateBook):
					res := &book{
						Name: m.textinput.Value(),
						Date: time.Now(),
					}
					appendEntry(f, res)
				case int(StateMovie):
					res := &movie{
						Name: m.textinput.Value(),
						Date: time.Now(),
					}
					appendEntry(f, res)
				case int(StateQuote):
					res := &quote{
						Quote: m.textinput.Value(),
						Date:  time.Now(),
					}
					appendEntry(f, res)
				case int(StateMusing):
					res := &musing{
						Musing: m.textinput.Value(),
						Date:   time.Now(),
					}
					appendEntry(f, res)
				}
				return m, tea.Quit
			}
		}
		m.textinput, cmd = m.textinput.Update(msg)
	}

	return m, cmd
}

func (m model) View() string {
	res := ""

	switch m.state {
	case int(StateStart):
		res += "What do you want to log?\n\n"
		for i, choice := range m.choices {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}

			res += fmt.Sprintf("%s %s\n", cursor, choice)
		}
	default:
		res += fmt.Sprintf("Log a %s\n", stateNames[m.state])
		res += m.textinput.View()
	}

	res += "\nPress ctrl+c to quit.\n"

	return res
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

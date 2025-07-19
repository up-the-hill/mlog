package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

const BASE_PATH = "/tmp/"

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
	textinput textinput.Model
}

func initialModel() model {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	return model{
		textinput: ti,
	}
}

func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("mlog")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit

		case "enter":
			appendEntry(BASE_PATH+"musings.ndjson", &musing{
				Musing: m.textinput.Value(),
				Date:   time.Now(),
			},
			)

			return m, tea.Quit
		}
	}
	m.textinput, cmd = m.textinput.Update(msg)

	return m, cmd
}

func (m model) View() string {
	res := ""

	res += "Log your thoughts\n"
	res += m.textinput.View()

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

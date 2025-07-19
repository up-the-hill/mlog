package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func initialModel() model {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 128
	ti.Width = 20
	return model{
		textinput: ti,
	}
}

func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("mlog")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.exiting {
		return m, tea.Quit
	}
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit

		case "enter":
			appendEntry(musingPath, &musing{
				Musing: m.textinput.Value(),
				Date:   time.Now(),
			},
			)
			m.exiting = true
			return m, nil
		}
	}
	m.textinput, cmd = m.textinput.Update(msg)

	return m, cmd
}

func (m model) View() string {
	res := ""

	if m.exiting {
		res += fmt.Sprintf("Logged \"%s\"", m.textinput.Value())
		return res
	}

	res += "Log your thoughts\n"
	res += m.textinput.View()

	res += "\nPress ctrl+c to quit.\n"

	return res
}

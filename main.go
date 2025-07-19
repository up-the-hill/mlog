package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

const musingPath = "/tmp/musings.ndjson"

type musing struct {
	Musing string    `json:"musing"`
	Date   time.Time `json:"date"`
}

type model struct {
	textinput textinput.Model
}

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
	listPtr := flag.Bool("l", false, "list all musings")
	flag.Parse()
	if *listPtr {
		for _, m := range getEntries(musingPath) {
			println(m.Musing)
		}
	} else {
		p := tea.NewProgram(initialModel())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	}
}

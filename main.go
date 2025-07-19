package main

import (
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const musingPath = "/tmp/musings.ndjson"

func main() {
	listPtr := flag.Bool("l", false, "list all musings")
	flag.Parse()
	if *listPtr {
		for _, m := range getEntries(musingPath) {
			dateStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
			fmt.Printf("%s %s\n", dateStyle.Render(m.Date.Format("2006-01-02")), m.Musing)
		}
	} else {
		p := tea.NewProgram(initialModel())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	}
}

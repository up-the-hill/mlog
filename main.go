package main

import (
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	config, err := loadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}
	musingPath := config.MusingsFile

	listPtr := flag.Bool("l", false, "list all musings")
	exportPtr := flag.Bool("x", false, "export all musings to markdown")
	flag.Parse()
	if *listPtr {
		// in list mode
		for _, m := range getEntries(musingPath) {
			dateStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
			fmt.Printf("%s %s\n", dateStyle.Render(m.Date.Format("2006-01-02")), m.Musing)
		}
	} else if *exportPtr {
		exportMusings(musingPath)
	} else {
		// run without any flags
		p := tea.NewProgram(initialModel(musingPath))
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	}
}

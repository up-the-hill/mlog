package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
)

type Musing struct {
	Musing string    `json:"musing"`
	Date   time.Time `json:"date"`
}

type Model struct {
	textinput   textinput.Model
	exiting     bool
	musingsPath string
}

func AppendEntry(filename string, m string) {
	newEntry := &Musing{
		Musing: m,
		Date:   time.Now(),
	}

	if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
		panic(err)
	}
	f, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	entryJson, _ := json.Marshal(newEntry)
	f.Write(append(entryJson, "\n"...))
}

func GetEntries(filename string) []Musing {
	dat, _ := os.ReadFile(filename)
	lines := strings.Split(string(dat), "\n")
	var entries []Musing

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue // skip empty lines
		}

		var m Musing
		if err := json.Unmarshal([]byte(line), &m); err != nil {
			// Skip malformed lines
			continue
		}
		entries = append(entries, m)
	}

	return entries
}

func ExportMusings(musingPath, exportPath string) {
	musings := GetEntries(musingPath)
	if err := os.MkdirAll(filepath.Dir(exportPath), 0755); err != nil {
		panic(err)
	}

	f, err := os.Create(exportPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("# Musings\n\n")

	for _, m := range musings {
		line := fmt.Sprintf("- %s: %s\n", m.Date.Format("2006-01-02"), m.Musing)
		f.WriteString(line)
	}
	fmt.Printf("Musings exported to %s\n", exportPath)
}

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func appendEntry(filename string, m string) {
	newEntry := &musing{
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

func getEntries(filename string) []musing {
	dat, _ := os.ReadFile(filename)
	lines := strings.Split(string(dat), "\n")
	var entries []musing

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue // skip empty lines
		}

		var m musing
		if err := json.Unmarshal([]byte(line), &m); err != nil {
			// Skip malformed lines
			continue
		}
		entries = append(entries, m)
	}

	return entries
}

func exportMusings(musingPath, exportPath string) {
	musings := getEntries(musingPath)
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

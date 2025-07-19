package main

import (
	"encoding/json"
	"os"
	"strings"
)

func appendEntry(filename string, newEntry any) {
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

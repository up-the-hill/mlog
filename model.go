package main

import (
	"time"

	"github.com/charmbracelet/bubbles/textinput"
)

type musing struct {
	Musing string    `json:"musing"`
	Date   time.Time `json:"date"`
}

type model struct {
	textinput textinput.Model
	exiting   bool
}

package main

import (
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model NodeGraph
type tickMsg time.Time

var text TextWidget
var container Container

func TuiStart() {

	p := tea.NewProgram(DefaultNodeGraph(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

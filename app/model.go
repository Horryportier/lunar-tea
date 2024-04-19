package main

import (
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	n "lunar-tea/node"
	w "lunar-tea/widgets"
)

type model NodeGraph
type tickMsg time.Time

var text w.TextWidget
var container w.Container
var list w.ListWidget

func DefaultNodeGraph() NodeGraph {

	header := text.New("20x40 grid forground color = math.Tan(x * y)", lipgloss.NewStyle().
		Bold(true))
	footer := text.New("made with bubbletea!", lipgloss.NewStyle().
		Foreground(lipgloss.Color("#6C50FF")).
		Bold(true).
		Border(lipgloss.RoundedBorder()),
	)
	color_square := container.New(
		n.VERTICAL,
		lipgloss.NewStyle().Align(lipgloss.Center),
		[]n.Node{header, containers[2], footer},
	)

	return NodeGraph{
		Graph: container.New(
			n.HORIZONTAL,
			lipgloss.NewStyle().Align(lipgloss.Center),
			[]n.Node{color_square,
				list.New(),
			},
		)}
}

func TuiStart() {

	p := tea.NewProgram(DefaultNodeGraph(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

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

func DefaultNodeGraph() NodeGraph {

	header := w.NewTextWidget("20x40 grid forground color = math.Tan(x * y)", lipgloss.NewStyle().
		Bold(true))
	footer := w.NewTextWidget("made with bubbletea!", lipgloss.NewStyle().
		Foreground(lipgloss.Color("#6C50FF")).
		Bold(true).
		Border(lipgloss.RoundedBorder()),
	)
	color_square := w.NewContainer(
		n.VERTICAL,
		lipgloss.NewStyle().Align(lipgloss.Center),
		[]n.Node{header, containers[2], footer},
	)

	return NodeGraph{
		Graph: w.NewContainer(
			n.HORIZONTAL,
			lipgloss.NewStyle().Align(lipgloss.Center),
			[]n.Node{color_square,
				w.NewListWidget([]string{
					"FOO", "BAR", "BAZ",
				}, w.ListStyle{}).Title("FOOBARBAZ").SetFilteringEnabled(true),
			},
		)}
}

func TuiStart() {

	p := tea.NewProgram(DefaultNodeGraph(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

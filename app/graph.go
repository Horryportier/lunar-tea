package main

import (
	"math"
	"strconv"

	n "lunar-tea/node"
	w "lunar-tea/widgets"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	current_container = 0
	containers        = []n.Node{container.New(n.VERTICAL, lipgloss.NewStyle(), MakeGrid(
		func(x int, y int) int { return int(math.Tan(float64(x) * float64(y))) })),
		container.New(n.VERTICAL, lipgloss.NewStyle(), MakeGrid(
			func(x int, y int) int { return x * y })),
		container.New(n.VERTICAL,
			lipgloss.NewStyle().Border(lipgloss.RoundedBorder()),
			[]n.Node{text.New("ligma balls", lipgloss.NewStyle())}),
	}
)

type NodeGraph struct {
	Graph n.Node
}
type calcXY func(x int, y int) int

func MakeGrid(fn calcXY) []n.Node {
	var cols []n.Node
	for x := 0; x < 20; x++ {
		var texts []n.Node
		for y := 0; y < 40; y++ {
			s := lipgloss.NewStyle().
				Foreground(lipgloss.
					Color(strconv.Itoa(fn(x, y))))
			texts = append(texts, text.New("█", s))

		}
		cols = append(cols, container.New(n.HORIZONTAL, lipgloss.NewStyle(), texts))
	}
	return cols
}

func (n NodeGraph) Init() tea.Cmd {
	return nil
}

func (ng NodeGraph) View() string {
	return ng.Graph.View()
}
func (ng NodeGraph) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return ng, tea.Quit
		case "tab":
			switch ng.Graph.(type) {
			case w.Container:
				if current_container == len(containers) {
					current_container = 0
				}
				ng.Graph.(w.Container).Nodes[0].(w.Container).Nodes[1] = containers[current_container]
				current_container++
			}
		}
	}

	m, cmd := ng.Graph.Update(message)
	switch m.(type) {
	case w.Container:
		ng.Graph = w.Container(m.(w.Container))
	}

	cmds = append(cmds, cmd)

	return ng, tea.Batch(cmds...)
}

package main

import (
	"math"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type NodeType int

type JointType int

const (
	NODE NodeType = iota
	LEAF

	VERTICAL JointType = iota
	HORIZONTAL
)

var (
	current_container = 0
	containers        = []Node{container.New(VERTICAL, lipgloss.NewStyle(), MakeGrid(
		func(x int, y int) int { return int(math.Tan(float64(x) * float64(y))) })),
		container.New(VERTICAL, lipgloss.NewStyle(), MakeGrid(
			func(x int, y int) int { return x * y })),
		container.New(VERTICAL,
			lipgloss.NewStyle().Border(lipgloss.RoundedBorder()),
			[]Node{text.New("ligma balls", lipgloss.NewStyle())}),
	}
)

type Node interface {
	Init() tea.Cmd
	Children() []Node
	View() string
	Type() NodeType
	Update(message tea.Msg) (tea.Model, tea.Cmd)
}

type NodeGraph struct {
	Graph Node
}
type calcXY func(x int, y int) int

func MakeGrid(fn calcXY) []Node {
	var cols []Node
	for x := 0; x < 20; x++ {
		var texts []Node
		for y := 0; y < 40; y++ {
			s := lipgloss.NewStyle().
				Foreground(lipgloss.
					Color(strconv.Itoa(fn(x, y))))
			texts = append(texts, text.New("█", s))

		}
		cols = append(cols, container.New(HORIZONTAL, lipgloss.NewStyle(), texts))
	}
	return cols
}

func DefaultNodeGraph() NodeGraph {

	header := text.New("20x40 grid forground color = math.Tan(x * y)", lipgloss.NewStyle().
		Bold(true))
	footer := text.New("made with bubbletea!", lipgloss.NewStyle().
		Foreground(lipgloss.Color("#6C50FF")).
		Bold(true).
		Border(lipgloss.RoundedBorder()),
	)

	return NodeGraph{
		Graph: Container{
			nodes: []Node{header, containers[2], footer},
			join:  VERTICAL,
			style: lipgloss.NewStyle().Align(lipgloss.Center),
		}}
}

func (n NodeGraph) Init() tea.Cmd {
	return nil
}

func (ng NodeGraph) View() string {
	return ng.Graph.View()
}
func (ng NodeGraph) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return ng, tea.Quit
		case "tab":
			switch ng.Graph.(type) {
			case Container:
				if current_container == len(containers) {
					current_container = 0
				}
				ng.Graph.(Container).nodes[1] = containers[current_container]
				current_container++
			}
		}
	}

	return ng, nil
}

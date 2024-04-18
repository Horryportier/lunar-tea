package main

import (
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

func DefaultNodeGraph() NodeGraph {
	c1 := container.New(VERTICAL, lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).Padding(2).
		Background(lipgloss.Color("#000003A")), []Node{
		text.New("DEEZ NUTS ON YOUR MOUTH", lipgloss.NewStyle()),
		text.New("DEEZ NUTS IN YOUR MOUTH",
			lipgloss.NewStyle().Foreground(lipgloss.Color("#FAF00FF")))},
	)
	c2 := container.New(VERTICAL, lipgloss.NewStyle(), []Node{
		text.New("DEEZ NUTS AROUND YOUR MOUTH", lipgloss.NewStyle()),
		text.New("DEEZ NUTS OUDSIDE YOUR MOUTH",
			lipgloss.NewStyle().Padding(2).Foreground(lipgloss.Color("#FA050FF")))},
	)
	return NodeGraph{
		Graph: Container{
			nodes: []Node{c1, c2},
			join:  HORIZONTAL,
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
		}
	}

	return ng, nil
}

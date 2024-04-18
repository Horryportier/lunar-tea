package main

import (
	tea "github.com/charmbracelet/bubbletea"
	w "lunar-test/widgets"
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

func (ng NodeGraph) Init() tea.Cmd {
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
			case w.Container:
				if current_container == len(containers) {
					current_container = 0
				}
				ng.Graph.(g.Container).nodes[1] = containers[current_container]
				current_container++
			}
		}
	}

	return ng, nil
}

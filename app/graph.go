package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type NodeType int

const (
	NODE NodeType = iota
	LEAF
)

type Node interface {
	Render() string
	Type() NodeType
	Update(message tea.Msg) (tea.Model, tea.Cmd)
}

type NodeGraph []struct {
	Graph Node
}

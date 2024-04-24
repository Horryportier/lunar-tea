package node

import (
	tea "github.com/charmbracelet/bubbletea"
)

type NodeType int

type JointType int

const (
	NODE NodeType = iota
	LEAF

	VERTICAL JointType = iota
	HORIZONTAL
)

func (j JointType) String() string {
	if j == VERTICAL {
		return "vertical"
	} else {
		return "horizontal"
	}
}

type Node interface {
	tea.Model
	Children() []Node
	Type() NodeType
	Marshal() ([]byte, error)
}

package node

import tea "github.com/charmbracelet/bubbletea"

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
	Marshal() ([]byte, error)
}

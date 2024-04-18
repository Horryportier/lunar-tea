package widgets

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	g "lunar-tea/graph"
)

type Container struct {
	nodes []g.Node
	join  g.JointType
	style lipgloss.Style
}

func (c Container) New(join g.JointType, style lipgloss.Style, nodes []g.Node) Container {
	return Container{
		nodes: nodes,
		join:  join,
		style: style,
	}
}

func (c Container) Init() tea.Cmd {
	return nil
}

func (c Container) Children() []g.Node {
	return c.nodes
}
func (c Container) View() string {
	var render string
	var rendered_children []string
	for _, child := range c.nodes {
		rendered_children = append(rendered_children, child.View())
	}

	if c.join == g.VERTICAL {
		for _, rendered_child := range rendered_children {
			render = lipgloss.JoinVertical(lipgloss.Center, render, rendered_child)
		}
	}
	if c.join == g.HORIZONTAL {
		for _, rendered_child := range rendered_children {
			render = lipgloss.JoinHorizontal(lipgloss.Center, render, rendered_child)
		}
	}

	return c.style.Render(render)
}
func (c Container) Type() g.NodeType {
	return g.NODE
}
func (c Container) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	return c, nil
}

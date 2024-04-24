package widgets

import (
	"encoding/json"
	n "lunar-tea/node"
	j "lunar-tea/serialize"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Container struct {
	Nodes []n.Node
	join  n.JointType
	style lipgloss.Style
}

func NewContainer(join n.JointType, style lipgloss.Style, nodes []n.Node) Container {
	return Container{
		Nodes: nodes,
		join:  join,
		style: style,
	}
}

func (c Container) Init() tea.Cmd {
	return nil
}

func (c Container) Children() []n.Node {
	return c.Nodes
}
func (c Container) View() string {
	var render string
	var rendered_children []string
	for _, child := range c.Nodes {
		rendered_children = append(rendered_children, child.View())
	}

	if c.join == n.VERTICAL {
		for _, rendered_child := range rendered_children {
			render = lipgloss.JoinVertical(lipgloss.Center, render, rendered_child)
		}
	}
	if c.join == n.HORIZONTAL {
		for _, rendered_child := range rendered_children {
			render = lipgloss.JoinHorizontal(lipgloss.Center, render, rendered_child)
		}
	}

	return c.style.Render(render)
}
func (c Container) Type() n.NodeType {
	return n.NODE
}
func (c Container) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	for i, child := range c.Nodes {
		m, cmd := child.Update(message)
		cmds = append(cmds, cmd)
		switch m.(type) {
		case TextWidget:
			c.Nodes[i] = TextWidget(m.(TextWidget))
		case ListWidget:
			c.Nodes[i] = ListWidget(m.(ListWidget))
		case Container:
			c.Nodes[i] = Container(m.(Container))
		default:
			panic("this souhld not happend")
		}
	}

	return c, tea.Batch(cmds...)
}

func (c Container) Marshal() ([]byte, error) {

	b, err := j.JsonMap(c, "container", func(T interface{}, m map[string]string) (map[string]string, error) {
		var node_json []string
		for _, node := range c.Nodes {
			b, err := node.Marshal()
			if err != nil {
				return m, err
			}
			node_json = append(node_json, string(b))
		}
		b, err := json.Marshal(node_json)
		if err != nil {
			return m, err
		}
		m["nodes"] = string(b)
		m["join"] = c.join.String()

		return m, nil
	})

	return b, err
}

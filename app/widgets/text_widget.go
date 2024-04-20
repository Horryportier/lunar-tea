package widgets

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	n "lunar-tea/node"
)

type TextWidget struct {
	s     string
	style lipgloss.Style
}

func NewTextWidget(t string, style lipgloss.Style) TextWidget {
	return TextWidget{s: t, style: style}
}

func (tw TextWidget) Init() tea.Cmd {
	return nil
}
func (tw TextWidget) Children() []n.Node {
	return []n.Node{}
}
func (tw TextWidget) View() string {
	return tw.style.Render(tw.s)
}
func (tw TextWidget) Type() n.NodeType {
	return n.LEAF
}
func (tw TextWidget) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	return tw, nil
}

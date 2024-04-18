package widgets

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	g "lunar-tea/graph"
)

type TextWidget struct {
	s     string
	style lipgloss.Style
}

func (tw TextWidget) New(t string, style lipgloss.Style) TextWidget {
	return TextWidget{s: t, style: style}
}

func (tw TextWidget) Init() tea.Cmd {
	return nil
}
func (tw TextWidget) Children() []g.Node {
	return []g.Node{}
}
func (tw TextWidget) View() string {
	return tw.style.Render(tw.s)
}
func (tw TextWidget) Type() g.NodeType {
	return g.LEAF
}
func (tw TextWidget) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	return tw, nil
}

package widgets

import (
	"encoding/json"
	n "lunar-tea/node"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TextWidget struct {
	s     string
	style lipgloss.Style
}

func NewTextWidget(t string, style lipgloss.Style) TextWidget {
	return TextWidget{s: t, style: style}
}

func (m TextWidget) Init() tea.Cmd {
	return nil
}
func (m TextWidget) Children() []n.Node {
	return []n.Node{}
}
func (m TextWidget) View() string {
	return m.style.Render(m.s)
}
func (m TextWidget) Type() n.NodeType {
	return n.LEAF
}
func (m TextWidget) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m TextWidget) Serialize() (string, string, error) {
	b, err := json.Marshal(m)
	return "text", string(b), err
}

func (m TextWidget) Marshal() ([]byte, error) {
	_map := make(map[string]string)
	_map["type"] = "text"

	b, err := json.Marshal(_map)
	return b, err
}

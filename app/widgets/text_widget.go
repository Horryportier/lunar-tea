package widgets

import (
	"encoding/json"
	n "lunar-tea/node"
	j "lunar-tea/serialize"
	s "lunar-tea/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TextWidget struct {
	s     string
	Style s.Style
}

func NewTextWidget(t string, style lipgloss.Style) TextWidget {
	return TextWidget{s: t, Style: s.Style(style)}
}

func (m TextWidget) Init() tea.Cmd {
	return nil
}
func (m TextWidget) Children() []n.Node {
	return []n.Node{}
}
func (m TextWidget) View() string {
	return m.Style.Into().Render(m.s)
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

func (t TextWidget) Marshal() ([]byte, error) {
	b, err := j.JsonMap(t, "text", func(T interface{}, m map[string]string) (map[string]string, error) {
		b, err := t.Style.Marshal()
		if err != nil {
			return m, err
		}
		m["style"] = string(b)

		return m, nil
	})

	return b, err
}

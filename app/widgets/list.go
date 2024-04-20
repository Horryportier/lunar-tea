package widgets

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	n "lunar-tea/node"
)

const listHeight = 14

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type ListStyle struct {
	Title           lipgloss.Style
	PaginationStyle lipgloss.Style
	HelpStyle       lipgloss.Style
}

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type ListWidget struct {
	list     list.Model
	choice   string
	quitting bool
}

func (m ListWidget) Title(s string) ListWidget {
	m.list.Title = s
	return m
}
func (m ListWidget) SetShowStatusBar(b bool) ListWidget {
	m.list.SetShowStatusBar(b)
	return m
}
func (m ListWidget) SetFilteringEnabled(b bool) ListWidget {
	m.list.SetFilteringEnabled(b)
	return m
}

func (m ListWidget) Init() tea.Cmd {
	return nil
}

func (m ListWidget) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = string(i)
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m ListWidget) View() string {
	if m.choice != "" {
		return quitTextStyle.Render(fmt.Sprintf("%s? Sounds good to me.", m.choice))
	}
	if m.quitting {
		return quitTextStyle.Render("Not hungry? That’s cool.")
	}
	return "\n" + m.list.View()
}

func intoItems[T ~string](l []T) []list.Item {
	var items []list.Item
	for _, i := range l {
		items = append(items, item(i))
	}
	return items
}

func NewListWidget(items []string, style ListStyle) ListWidget {

	list_items := intoItems(items)
	const defaultWidth = 20

	l := list.New(list_items, itemDelegate{}, defaultWidth, listHeight)
	l.Styles.Title = style.Title
	l.Styles.PaginationStyle = style.PaginationStyle
	l.Styles.HelpStyle = style.HelpStyle

	return ListWidget{list: l}
}
func (m ListWidget) Children() []n.Node {
	return []n.Node{}
}

func (m ListWidget) Type() n.NodeType {
	return n.LEAF
}

package style

import (
	j "lunar-tea/serialize"
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

type Style lipgloss.Style

type color struct {
	r, g, b, a uint32
}

func fromTerminalColor(tc lipgloss.TerminalColor) color {
	r, g, b, a := tc.RGBA()
	return color{
		r, g, b, a,
	}
}

func (c color) Marshal() ([]byte, error) {
	buf, err := j.JsonMap(c, "color", func(T interface{}, m map[string]string) (map[string]string, error) {
		m["a"] = strconv.Itoa(int(c.a))
		m["r"] = strconv.Itoa(int(c.r))
		m["g"] = strconv.Itoa(int(c.g))
		m["b"] = strconv.Itoa(int(c.b))
		return m, nil
	})

	return buf, err
}

func (s Style) Marshal() ([]byte, error) {
	buf, err := j.JsonMap(s, "style", func(T interface{}, m map[string]string) (map[string]string, error) {
		style := lipgloss.Style(T.(Style))
		bg, err := fromTerminalColor(style.GetBackground()).Marshal()
		m["background"] = string(bg)
		return m, err
	})
	return buf, err
}

func (s Style) Into() lipgloss.Style {
	return lipgloss.Style(s)
}

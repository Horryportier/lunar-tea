package main

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
	buf, err := j.JsonMap(c, func(T interface{}, m map[string]string) (map[string]string, error) {
		color := color(T.(color))
		m["type"] = "color"
		m["a"] = strconv.Itoa(int(color.a))
		m["r"] = strconv.Itoa(int(color.r))
		m["g"] = strconv.Itoa(int(color.g))
		m["b"] = strconv.Itoa(int(color.b))
		return m, nil
	})
	return buf, err
}

func (s Style) Marshal() ([]byte, error) {
	buf, err := j.JsonMap(s, func(T interface{}, m map[string]string) (map[string]string, error) {
		style := lipgloss.Style(T.(Style))
		m["type"] = "style"
		bg, err := fromTerminalColor(style.GetBackground()).Marshal()
		m["background"] = string(bg)
		return m, err
	})
	return buf, err
}

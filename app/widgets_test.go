package main

import (
	"fmt"
	w "lunar-tea/widgets"
	"testing"

	"github.com/charmbracelet/lipgloss"
)

func TestSerialize(t *testing.T) {
	list := w.NewListWidget([]string{"FOO", "BAR"}, w.ListStyle{})
	b, e := list.Marshal()
	if e != nil {
		t.Fatal(e)
	}

	fmt.Printf("json: %s\n", string(b))
	text := w.NewTextWidget("bar", lipgloss.NewStyle())
	b, e = text.Marshal()
	if e != nil {
		t.Fatal(e)
	}
	fmt.Printf("json: %s\n", string(b))

	b, e = Style(lipgloss.NewStyle().Background(lipgloss.Color("23"))).Marshal()
	if e != nil {
		t.Fatal(e)
	}
	fmt.Printf("json: %s\n", string(b))

}

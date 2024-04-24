package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestSerialize(t *testing.T) {
	b, err := DefaultNodeGraph().Marshal()
	if err != nil {
		t.Fatal(err)
	}
	json_string := string(b)
	for i := 0; i <= 10; i++ {
	 json_string =	strings.ReplaceAll(json_string, "\\\\", "\\")
	}
	fmt.Println(json_string)
	//	list := w.NewListWidget([]string{"FOO", "BAR"}, w.ListStyle{})
	//	b, e := list.Marshal()
	//	if e != nil {
	//		t.Fatal(e)
	//	}
	//
	//	fmt.Printf("json: %s\n", string(b))
	//	text := w.NewTextWidget("bar", lipgloss.NewStyle())
	//	b, e = text.Marshal()
	//	if e != nil {
	//		t.Fatal(e)
	//	}
	//	fmt.Printf("json: %s\n", string(b))
	//
	//	b, e = Style(lipgloss.NewStyle().Background(lipgloss.Color("23"))).Marshal()
	//	if e != nil {
	//		t.Fatal(e)
	//	}
	//	fmt.Printf("json: %s\n", string(b))

}

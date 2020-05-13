package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConvertToImageLayers(t *testing.T) {
	tests := []struct {
		input  string
		width  int
		height int
		output [][][]int
	}{
		{"123456789012", 3, 2, [][][]int{{{1, 2, 3}, {4, 5, 6}}, {{7, 8, 9}, {0, 1, 2}}}},
	}

	for _, test := range tests {
		got := convertToImageLayers(test.input, test.width, test.height)
		if !reflect.DeepEqual(got, test.output) {
			t.Errorf("got %v want %v", got, test.output)
		}
		fmt.Printf("%v", len(got))
	}
}

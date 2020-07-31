package main

import (
	"reflect"
	"testing"
)

func TestConvertStrArrayToIntArray(t *testing.T) {
	input := []string{"1", "2", "999"}
	want := []int{1, 2, 999}

	got := convertStrArrayToIntArray(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

// Additional tests from https://www.reddit.com/r/adventofcode/comments/e8aw9j/2019_day_9_part_1_how_to_fix_203_error/fac3294/
func TestIntcodeMachine(t *testing.T) {
	opcodeTests := []struct {
		program []int
		input   int
		output  []int
	}{
		// {[]int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}, 1, []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}},
		// {[]int{1102, 34915192, 34915192, 7, 4, 7, 99, 0}, 1, []int{1219070632396864}},
		// {[]int{104, 1125899906842624, 99}, 1, []int{1125899906842624}},
		// {[]int{109, -1, 4, 1, 99}, 1, []int{-1}},
		// {[]int{109, -1, 104, 1, 99}, 1, []int{1}},
		// {[]int{109, -1, 204, 1, 99}, 1, []int{109}},
		// {[]int{109, 1, 9, 2, 204, -6, 99}, 1, []int{204}},
		// {[]int{109, 1, 109, 9, 204, -6, 99}, 1, []int{204}},
		// {[]int{109, 1, 209, -1, 204, -106, 99}, 1, []int{204}},
		// {[]int{109, 1, 3, 3, 204, 2, 99}, 444, []int{444}},
		{[]int{109, 1, 203, 2, 204, 2, 99}, 666, []int{666}},
	}

	for _, test := range opcodeTests {
		got := checkIntcodeMachine(test.program, test.input, true)
		if !reflect.DeepEqual(got, test.output) {
			t.Errorf("got %v want %v", got, test.output)
		}
	}
}

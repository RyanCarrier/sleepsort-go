package sleepsortgo

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

var inputs = [][]int{
	{3, 2, 1},
	{1, 2, 3},
	{1, 1, 1},
	{1, 2, 1},
	{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
}

func TestSleepSort(t *testing.T) {
	for i, input := range inputs {
		output := SleepSort(input, time.Millisecond)
		sort.Ints(input)
		if !reflect.DeepEqual(output, input) {
			t.Errorf("SleepSort(%v) = %v; want %v", inputs[i], output, input)
		}
	}
}

func TestSleepSortDefault(t *testing.T) {
	for i, input := range inputs {
		output := SleepSortDefault(input)
		sort.Ints(input)
		if !reflect.DeepEqual(output, input) {
			t.Errorf("SleepSort(%v) = %v; want %v", inputs[i], output, input)
		}
	}
}

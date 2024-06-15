package unit4

import (
	"testing"
)

func TestRecursiveSum(t *testing.T) {
	var tests = []struct {
		name string
		args []int
		want int
	}{
		{"1+5", []int{1, 5}, 6},
		{"4+6+4+2", []int{4, 6, 4, 2}, 16},
		{"noneData", []int{}, 0},
		{"oneInt", []int{10}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := recursiveSum(tt.args)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestRecursiveMax(t *testing.T) {
	var tests = []struct {
		name string
		args []int
		want int
	}{
		{"1+5", []int{1, 5}, 5},
		{"4+6+4+2", []int{4, 6, 4, 2}, 6},
		{"noneData", []int{}, 0},
		{"oneInt", []int{10}, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := recursiveMax(tt.args)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	var tests = []struct {
		name string
		args []int
		want []int
	}{
		{"6,1,3,7,3", []int{6, 1, 3, 7, 3}, []int{1, 3, 3, 6, 7}},
		{"5, 2, 6, 2, 5, 6, 2, 3, 4, 5, 1, 6, 7, 1, 43, 1", []int{5, 2, 6, 2, 5, 6, 2, 3, 4, 5, 1, 6, 7, 1, 43, 1}, []int{1, 1, 1, 2, 2, 2, 3, 4, 5, 5, 5, 6, 6, 6, 7, 43}},
		{"noneData", []int{}, []int{}},
		{"oneInt", []int{10}, []int{10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := quickSort(tt.args)
			if len(ans) != len(tt.want) {
				t.Errorf("got len(ans)=%d, want %d", len(ans), len(tt.want))
				t.Fail()
				return
			}
			for i, v := range ans {
				if v != tt.want[i] {
					t.Errorf("got ans[%d] = %d, want %d", i, v, tt.want[i])
					t.Fail()
					return
				}
			}
		})
	}

}

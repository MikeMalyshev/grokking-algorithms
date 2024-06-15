package unit1

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	data := make([]int, 100)
	for i := range data {
		data[i] = i
	}
	ans := binarySearch(53, data)
	if ans != 53 {
		t.Errorf("binarySearch(53,data{0...100}) = %d; want 53", ans)
	}
}

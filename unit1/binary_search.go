package unit1

func binarySearch(value int, data []int) int {
	midIdx := len(data) / 2

	if value == data[midIdx] {
		return midIdx
	}
	if len(data) == 1 {
		return -1
	}
	if value > data[midIdx] {
		return midIdx + binarySearch(value, data[midIdx:])
	}
	return binarySearch(value, data[:midIdx])
}

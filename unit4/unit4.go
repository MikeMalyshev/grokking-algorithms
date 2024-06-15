package unit4

func recursiveSum(data []int) int {
	if len(data) < 1 {
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}
	return data[0] + recursiveSum(data[1:])
}

func recursiveMax(data []int) int {
	if len(data) < 1 {
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}

	max := recursiveMax(data[1:])
	if data[0] > max {
		return data[0]
	}
	return max
}

func quickSort(data []int) []int {
	if len(data) < 2 {
		return data
	}

	cp := data[0]
	var smallerValues, largerValues []int
	for _, v := range data[1:] {
		if cp > v {
			smallerValues = append(smallerValues, v)
			continue
		}
		largerValues = append(largerValues, v)
	}
	smallerValues = quickSort(smallerValues)
	largerValues = quickSort(largerValues)

	return append(append(smallerValues, cp), largerValues...)
}

package main

func binarySearch(values []int, target int, lowIdx int, highIdx int) int {
	if highIdx < lowIdx || len(values) == 0 {
		return -1
	}
	middle := (highIdx + lowIdx) / 2
	if values[middle] > target {
		return binarySearch(values, target, lowIdx, middle-1)
	} else if values[middle] < target {
		return binarySearch(values, target, middle+1, highIdx)
	} else {
		return middle
	}
}

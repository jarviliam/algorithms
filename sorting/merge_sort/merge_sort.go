/**
* Merge Sort Golang
*/
package main

import (
	"fmt"
)

/**
* Merges the left and right depending on whichever is smaller
* If there is no more entries in either left or right, append the rest
*/
func merge(left, right []int) []int {
	output := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(output, right...)
		}
		if len(right) == 0 {
			return append(output, left...)
		}
		if left[0] <= right[0] {
			output = append(output, left[0])
			left = left[1:]
		} else {
			output = append(output, right[0])
			right = right[1:]
		}
	}
	return output
}

/**
* If Array has only one entry return
* Otherwise Recursion split it then merge
*/
func mergeSort(in []int) []int {
	if len(in) <= 1 {
		return in
	}
	middle := len(in) / 2
	left := mergeSort(in[:middle])
	right := mergeSort(in[middle:])
	return merge(left, right)
}

func main() {
	i := []int{1, 44, 22, 92, 94, 12, 3323, 52}
    out := mergeSort(i)
	for _, v := range out {
		fmt.Printf("Value: %d\n", v)
	}
}

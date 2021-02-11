/**
* https://www.hackerearth.com/practice/algorithms/searching/binary-search/practice-problems/algorithm/do-you-really-understand-binary-search/
*
*
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sliceAtoi(a []string) ([]int, error) {
	sa := make([]int, 0, len(a))
	for _, v := range a {
		i, error := strconv.Atoi(v)
		if error != nil {
			return sa, error
		}
		sa = append(sa, i)
	}
	return sa, nil
}

func load(in string) ([]int, []int, error) {
	file, err := os.Open(in)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()
	var input []int
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	//fmt.Println(scanner.Text())
	scanner.Scan()
	temp := strings.Split(scanner.Text(), " ")
	converted, _ := sliceAtoi(temp)
	scanner.Scan()
	numInputs, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < numInputs; i++ {
		scanner.Scan()
		v, _ := strconv.Atoi(strings.Split(scanner.Text(), " ")[0])
		input = append(input, v)
	}
	return input, converted, scanner.Err()
}
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

func main() {
	points, lines, _ := load("ex1.txt")
	for _, v := range points {
		out := binarySearch(lines, v, 0, len(lines)-1)
		if out == -1 {
			fmt.Println(0)
		} else {
			fmt.Println(out)
		}
	}
}

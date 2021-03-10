package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
    "sort"
	"strings"
)

func conv(in []string) []int {
	var out []int
	for _, v := range in {
		x, _ := strconv.Atoi(v)
		out = append(out, x)
	}
	return out
}


func main() {
	file, _ := os.Open("MMM1.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	values := conv(strings.Split(scanner.Text(), " "))
    sort.Ints(values)

	//Mean
	mean := float64(0)
	for _, v := range values {
		mean += float64(v)
	}
	mean = (float64(mean) / float64(num))

	//Median
	median := 0
	if num%2 == 0 {
		median = (values[num/2] + values[num/2] - 1) / 2
	} else {
		median = values[(num-1)/2]
	}

	//Mode
	count := make(map[int]int)
	for _, v := range values {
		count[v]++
	}

	mode := 0
	max := 0
	for _, v := range values {
		if count[v] > max {
			mode = v
			max = count[v]
		}
	}
	fmt.Println("Mode ", mode)
	fmt.Println("Median ", median)
	fmt.Println("Mean ", mean)

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	file, _ := os.Open("weight1.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	scanner.Scan()
	values := conv(strings.Split(scanner.Text(), " "))
	scanner.Scan()
	idx := conv(strings.Split(scanner.Text(), " "))

	mean := 0

	for i, v := range values {
		mean = mean + v*idx[i]
	}
	denom := 0
	for _, v := range idx {
		denom = denom + v
	}
	var answer float64
	answer = float64(mean) / float64(denom)
	fmt.Printf(strconv.FormatFloat(answer, 'f', -1, 64))
}

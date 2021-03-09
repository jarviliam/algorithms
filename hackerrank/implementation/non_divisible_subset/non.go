/**
* Non-Divisible Subset
* @url - https://www.hackerrank.com/challenges/non-divisible-subset/problem
* Difficulty - Medium
* Time : 15mins
* Explanation :
*  Make A counting Array filled with results from x % Divisor
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

func toInt(in []string) []int {
	var out []int
	for _, v := range in {
		x, _ := strconv.Atoi(v)
		out = append(out, x)
	}
	return out
}

func prepare(input string) ([]int, int, error) {
	file, err := os.Open(input)
	if err != nil {
		return nil, -1, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line1 := toInt(strings.Split(scanner.Text(), " "))
	scanner.Scan()
	line2 := toInt(strings.Split(scanner.Text(), " "))
	return line2, line1[1], nil
}

func main() {

	vals, divisor, _ := prepare("in2.txt")
	remainders := make([]int, divisor+1)
	for _, v := range vals {
		r := (v % divisor)
		remainders[r]++
	}
	ans := 0
	if remainders[0] > 0 {
		ans++
	}
	for i, v := range remainders {
		if i == (divisor - i) {
			ans++
		} else {
			if v >= remainders[divisor-i] {
				ans += v
			} else {
				ans += remainders[divisor-i]
			}
			remainders[i] = 0
			remainders[divisor-i] = 0
		}
	}
	fmt.Println("Answer ", ans)
}

/*
 - https://www.hackerrank.com/challenges/icecream-parlor/problem
 - Although the Problem State's Binary Search, It's a lot easier
 - to just use a Hash Map
 - Easy
 - 2021-02-15
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ProblemSet struct {
	money    int
	flavours []int
}

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

func load(in string) ([]ProblemSet, error) {
	file, err := os.Open(in)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	sets, _ := strconv.Atoi(scanner.Text())
	var input []ProblemSet
	for i := 0; i < sets; i++ {
		scanner.Scan()
		money, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		scanner.Scan()
		prices, _ := sliceAtoi(strings.Split(scanner.Text(), " "))
		input = append(input, ProblemSet{money, prices})
	}
	return input, scanner.Err()
}

func solve(problem ProblemSet) {
	var check map[int][]int
	check = make(map[int][]int)
	for i, v := range problem.flavours {
		if _, ok := check[v]; ok {
			check[v] = append(check[v], i)
		} else {
			check[v] = []int{i}
		}
	}
	for i, v := range problem.flavours {
		target := problem.money - v
		if value, ok := check[target]; ok {
			if value[0] != i {
				fmt.Println("Answer ", check[target][0]+1, i+1)
				break
			}
		}
	}
}

func main() {
	problems, _ := load("icecream.txt")
	for _, v := range problems {
		solve(v)
	}
}

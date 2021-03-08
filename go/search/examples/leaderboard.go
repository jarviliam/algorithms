/*
* Climbing The LeaderBoard
* https://www.hackerrank.com/challenges/climbing-the-leaderboard/problem#!
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

type ProblemSet struct {
	leaderboard []int
	scores      []int
}

func sliceAtoi(in []string) ([]int, error) {
	var out []int
	for _, v := range in {
		o, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		out = append(out, o)
	}
	return out, nil
}

func load(path string) (*ProblemSet, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	scanner.Scan()
	leader, _ := sliceAtoi(strings.Split(scanner.Text(), " "))
	scanner.Scan()
	scanner.Scan()
	score, _ := sliceAtoi(strings.Split(scanner.Text(), " "))
	return &ProblemSet{clearDupes(leader), score}, nil
}
func clearDupes(in []int) []int {
	var out []int
	seen := make(map[int]struct{})
	for _, v := range in {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			out = append(out, v)
		}
	}
	return out
}

func binary(set []int, target int, high int, low int) int {
	if high < low {
		if set[low] > target {
			return low
		}
		if set[low] < target {
			return low + 1
		}
		fmt.Println(high, low)
		return -1
	}
	center := (high + low) / 2
	if set[center] > target {
		fmt.Println("Split bottom", target, high, center+1)
		return binary(set, target, center-1, low)
	}
	if set[center] < target {
		fmt.Println("Split top", target, center-1, low)
		return binary(set, target, high, center+1)
	}
	//Otherwise Pos is Equal
	return center
}

func main() {
	problem, _ := load("leaderboard.txt")
	ans := binary(problem.leaderboard, problem.scores[0], 0,len(problem.leaderboard)-1)
	fmt.Println("Ans ", ans)

	/*for _, v := range problem.scores {
	        fmt.Println("B - ",v)
			ans := binary(problem.leaderboard, v, len(problem.leaderboard)-1, 0)
			fmt.Println(ans)
		}*/
}

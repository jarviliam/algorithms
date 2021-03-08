/*
   https://www.hackerrank.com/challenges/gridland-metro/problem
   Difficulty - Medium (Although it really isnt)
   2021-02-16
   Time: 5 mins
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
	size   int
	tracks []Track
}

type Track struct {
	row   int
	start int
	end   int
}

func sliceAtoi(in []string) ([]int, error) {
	var out []int
	for _, v := range in {
		ex, error := strconv.Atoi(v)
		if error != nil {
			return nil, error
		}
		out = append(out, ex)
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
	first, _ := sliceAtoi(strings.Split(scanner.Text(), " "))
	//Get Size of Grid
	size := first[1] * first[0]

	var tracks []Track

	for i := 0; i < first[2]; i++ {
		scanner.Scan()
		values, _ := sliceAtoi(strings.Split(scanner.Text(), " "))
		tracks = append(tracks, Track{values[0], values[1], values[2]})
	}
	return &ProblemSet{size, tracks}, nil
}

type Pair struct {
	start, end int
}

func overlappingCase(problem *ProblemSet) int {
	size := problem.size
	queue := make(map[int]Pair)
	for _, v := range problem.tracks {
		if value, ok := queue[v.row]; ok {
			if value.start > v.start {
				queue[v.row] = Pair{v.start, value.end}
			}
			if value.end < v.end {
				queue[v.row] = Pair{value.start, v.end}
			}
		} else {
			queue[v.row] = Pair{v.start, v.end}
		}
	}
	//Calculate
	for _, v := range queue {
		size -= v.end - v.start + 1
	}
	return size
}

func main() {
	fmt.Println("Lets go")
	problemSet, _ := load("grid.txt")
	//This is If We Do Not Account for Overlapping Rows
	for _, v := range problemSet.tracks {
		problemSet.size -= (v.end - v.start) + 1
	}
	fmt.Println(problemSet.size)
	//This is For the Overlapping rows
	fmt.Println(overlappingCase(problemSet))
}

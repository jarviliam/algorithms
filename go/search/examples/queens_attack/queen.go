/**
* Queens Attack 2 - HackerRank Problem
* @url - https://www.hackerrank.com/challenges/queens-attack-2/problem
* Difficulty - Medium
* Time : 20mins
* Explanation :
* 1) Record How Many Spaces the Queen can Move to the Edge in Each Cardinal Direction
* 2) At Every Obstacle, Determine What cardinal Direction it intercepts
* 3) Basically like a HashMap with Cardinal Directions, if the current obstacle is closer
* than another obstacle in the same cardinal direction, we update it as the closest
* 4) from there sum the values of each direction
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	x int
	y int
}

type Problem struct {
	size      int
	qX        int
	qY        int
	obstacles []Coord
}

func arrConv(inp []string) []int {
	var out []int
	for _, v := range inp {
		z, _ := strconv.Atoi(v)
		out = append(out, z)
	}
	return out
}

func parse(path string) (*Problem, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line1 := arrConv(strings.Split(scanner.Text(), " "))
	scanner.Scan()
	line2 := arrConv(strings.Split(scanner.Text(), " "))
	var coords []Coord
	for scanner.Scan() {
		v := arrConv(strings.Split(scanner.Text(), " "))
		coords = append(coords, Coord{v[0], v[1]})
	}
	return &Problem{line1[0], line2[0], line2[1], coords}, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func main() {
	problem, _ := parse("in2.txt")

	var directions [8]int
	directions[0] = min(problem.qX, problem.size+1-problem.qY)                // Top Left
	directions[1] = problem.size + 1 - problem.qY                             // top
	directions[2] = min(problem.size+1-problem.qX, problem.size+1-problem.qY) //Top right
	directions[3] = problem.qX                                                // Left
	directions[4] = problem.size + 1 - problem.qX                             // Right
	directions[5] = min(problem.qX, problem.qY)                               // Bottom Left
	directions[6] = problem.qY                                                // Bottom
	directions[7] = min(problem.size+1-problem.qX, problem.qY)                //Bottom Right

	for _, v := range problem.obstacles {
		dx := v.x - problem.qX
		dy := v.y - problem.qY
		absY := abs(dy)
		absX := abs(dx)
		if dx == 0 {
			//up Or down
			//Down
			if dy > 0 {
				if directions[6] > absY {
					directions[6] = absY
				}
			} else {
				//up
				if directions[1] > absY {
					directions[1] = absY
				}
			}
		} else if dy == 0 {
			//Left Rigth
			if dx < 0 {
				if directions[3] > absX {
					directions[3] = absX
				}
			} else {
				if directions[4] > absX {
					directions[4] = absX
				}
			}
		} else if dx == dy {
			//Diagnols
			if dx < 0 {
				//Top left
				if dy < 0 {
					if directions[0] > absX {
						directions[0] = absX
					}
				} else {
					//Bottom Left
					if directions[5] > absX {
						directions[5] = absX
					}
				}
			} else {
				if dy < 0 {
					//Top Right
					if directions[2] > absX {
						directions[2] = absX
					}
				} else {
					if directions[7] > absX {
						directions[7] = absX
					}
				}
			}
		} else {
			continue
		}
	}
	final := 0
	for _, v := range directions {
		final += v - 1
	}
	fmt.Println("Final Total ", final)
}

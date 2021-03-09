/**
* Extra Long Factorials
* @url - https://www.hackerrank.com/challenges/extra-long-factorials/problem
* Difficulty - Medium
* Time : 10mins
* Explanation :
* Very simple, Just Recursivly call to 0
 */

package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func factorial(n int) *big.Int {

	x := big.NewInt(int64(n))
	zr := big.NewInt(int64(0))
	if x.Cmp(zr) == 0 {
		return big.NewInt(int64(1))
	} else {
		return x.Mul(x, factorial(n-1))
	}
}

func main() {
	num, _ := strconv.Atoi(os.Args[1:][0])
	fmt.Println("answer ", factorial(num))
}

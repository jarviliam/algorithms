/**
*
*
*
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func load(in string) ([]string, error) {
	file, err := os.Open(in)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var input []string
	scanner := bufio.NewScanner(file)
	numLines := strconv.Atoi(scanner.Text())
	/*for scanner.Scan() {
		input = append(input, scanner.Text())
	}*/
	fmt.Println(numLines)
	return input, scanner.Err()
}

func main() {

}

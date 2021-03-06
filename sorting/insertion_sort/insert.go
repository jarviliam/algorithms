package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	array, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(errors.New("Error Input"))
	}
	trimmed := strings.Split(strings.TrimSpace(array), " ")

	var convArrayA []int
	for _, v := range trimmed {
		x, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		convArrayA = append(convArrayA, x)
	}
	start := time.Now()
	fmt.Println("Sorted : ", InsertionSort(convArrayA))
	end := time.Now()
	fmt.Println("Running time is ", end.Sub(start))
}

func InsertionSort(in []int) []int {
	if len(in) == 0 {
		return nil
	} else if len(in) == 1 {
		return in
	} else {
		for i := 0; i < len(in); i++ {
			j := i
			v := in[j]

			for i > 0 && v < in[i-1] {
				in[i] = in[i-1]
				i--
			}
			in[i] = v
		}
	}
	return in
}

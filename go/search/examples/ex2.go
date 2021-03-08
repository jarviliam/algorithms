package main
/**
*https://www.hackerearth.com/practice/algorithms/searching/binary-search/practice-problems/algorithm/when-will-she-talk/
* 2021-02-13
*
*/

import(
    "fmt"
    "bufio"
    "os"
    "strconv"
    "sort"
    "strings"
)

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

func load(in string) ([]int, []int, error) {
	file, err := os.Open(in)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()
	var input []int
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	scanner.Scan()
	temp := strings.Split(scanner.Text(), " ")
	converted, _ := sliceAtoi(temp)
	scanner.Scan()
	numInputs, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < numInputs; i++ {
		scanner.Scan()
		v, _ := strconv.Atoi(strings.Split(scanner.Text(), " ")[0])
		input = append(input, v)
	}
	sort.Slice(converted, func(i, j int) bool {
		return converted[i] < converted[j]
	})
	return input, converted, scanner.Err()
}

func main(){
    fmt.Println("Example 2")

    days,queries,_ := load("2.txt")

}

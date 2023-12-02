package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func isNumber(r byte) (bool, int) {
	num, err := strconv.Atoi(string(r))
	if err == nil {
		return true, num
	}
	return false, 0
}

func main() {
	f, err := os.Open("../data.txt")
	check(err)
	result := [][]int{}
	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() {
		digits := []int{-1, -1}
		line := fileScanner.Text()
		for i := 0; i < len(line); i++ {
			if digits[0] != -1 && digits[1] != -1 {
				break
			}
			if digits[0] == -1 {
				isNum_first, num_first := isNumber(line[i])
				if isNum_first {
					digits[0] = num_first
				}
			}
			if digits[1] == -1 {
				isNum_last, num_last := isNumber(line[len(line)-i-1])
				if isNum_last {
					digits[1] = num_last
				}
			}
		}
		result = append(result, digits)
	}
	sum := 0
	for _, row := range result {
		sum += row[0]*10 + row[1]
	}
	fmt.Println(sum)

	f.Close()
}


package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("day1_test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var totalSum int

	for scanner.Scan() {
		line := scanner.Text()
		calibVal := getCalibrationValue(line)
		totalSum += calibVal
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(totalSum)
	}
}

func getCalibrationValue(line string) int {
	left, right := 0, len(line)-1

	for left < right {
		if !unicode.IsDigit(rune(line[left])) {
			left++
			continue
		}
		if !unicode.IsDigit(rune(line[right])) {
			right--
			continue
		}
		break
	}

	var calibValStr string
	if left == right {
		calibValStr = string(line[left]) + string(line[left])
	} else {
		calibValStr = string(line[left]) + string(line[right])
	}

	res, _ := strconv.Atoi(calibValStr)
	return res
}

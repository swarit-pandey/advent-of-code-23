package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

// Messed up the second part lol
func main() {
	file, err := os.Open("day3_test.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var matrix [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matrix = append(matrix, []rune(scanner.Text()))
	}

	sum := 0
	processed := make([][]bool, len(matrix))
	for i := range processed {
		processed[i] = make([]bool, len(matrix[i]))
	}

	for y, row := range matrix {
		for x, char := range row {
			if unicode.IsDigit(char) && isAdjacentToSym(matrix, x, y) && !processed[y][x] {
				number := getFullNumber(matrix, x, y, &processed)
				if num, err := strconv.Atoi(number); err == nil {
					sum += num
				}
			}
		}
	}

	fmt.Println(sum)

	part2()
}

func isAdjacentToSym(matrix [][]rune, x, y int) bool {
	dirs := []struct{ dx, dy int }{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for _, d := range dirs {
		nx, ny := x+d.dx, y+d.dy
		if nx >= 0 && ny >= 0 && ny < len(matrix) && nx < len(matrix[ny]) && !unicode.IsDigit(matrix[ny][nx]) && matrix[ny][nx] != '.' {
			return true
		}
	}
	return false
}

func getFullNumber(matrix [][]rune, x, y int, processed *[][]bool) string {
	number := string(matrix[y][x])
	(*processed)[y][x] = true
	for i := x - 1; i >= 0 && unicode.IsDigit(matrix[y][i]) && !(*processed)[y][i]; i-- {
		number = string(matrix[y][i]) + number
		(*processed)[y][i] = true
	}
	for i := x + 1; i < len(matrix[y]) && unicode.IsDigit(matrix[y][i]) && !(*processed)[y][i]; i++ {
		number += string(matrix[y][i])
		(*processed)[y][i] = true
	}
	for j := y - 1; j >= 0 && unicode.IsDigit(matrix[j][x]) && !(*processed)[j][x]; j-- {
		number = string(matrix[j][x]) + number
		(*processed)[j][x] = true
	}
	for j := y + 1; j < len(matrix) && unicode.IsDigit(matrix[j][x]) && !(*processed)[j][x]; j++ {
		number += string(matrix[j][x])
		(*processed)[j][x] = true
	}
	return number
}

func part2() {
	file, err := os.Open("day3_example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var matrix [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matrix = append(matrix, []rune(scanner.Text()))
	}

	gearSum := 0
	for y, row := range matrix {
		for x, char := range row {
			if char == '*' {
				ratios := getGearRatios(matrix, x, y)
				if len(ratios) == 2 {
					gearSum += ratios[0] * ratios[1]
				}
			}
		}
	}

	fmt.Println(gearSum)
}

func getGearRatios(matrix [][]rune, x, y int) []int {
	dirs := []struct{ dx, dy int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var ratios []int
	for _, d := range dirs {
		nx, ny := x+d.dx, y+d.dy
		if nx >= 0 && ny >= 0 && ny < len(matrix) && nx < len(matrix[ny]) && unicode.IsDigit(matrix[ny][nx]) {
			number := getFullNumber2(matrix, nx, ny)
			if num, err := strconv.Atoi(number); err == nil {
				ratios = append(ratios, num)
			}
		}
	}
	return ratios
}

func getFullNumber2(matrix [][]rune, x, y int) string {
	number := string(matrix[y][x])
	for i := x - 1; i >= 0 && unicode.IsDigit(matrix[y][i]); i-- {
		number = string(matrix[y][i]) + number
	}
	for i := x + 1; i < len(matrix[y]) && unicode.IsDigit(matrix[y][i]); i++ {
		number += string(matrix[y][i])
	}
	for j := y - 1; j >= 0 && unicode.IsDigit(matrix[j][x]); j-- {
		number = string(matrix[j][x]) + number
	}
	for j := y + 1; j < len(matrix) && unicode.IsDigit(matrix[j][x]); j++ {
		number += string(matrix[j][x])
	}
	return number
}

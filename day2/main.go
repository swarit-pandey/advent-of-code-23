package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day2_test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var validGames []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if isValidGame(line) {
			gameID := getGameID(line)
			validGames = append(validGames, gameID)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sum(validGames))
	}
}

func isValidGame(line string) bool {
	parts := strings.Split(line, ": ")
	subsets := strings.Split(parts[1], "; ")
	maxCubes := map[string]int{"red": 0, "green": 0, "blue": 0}

	for _, subset := range subsets {
		currentCubes := parseCubes(subset)
		for color, count := range currentCubes {
			if count > maxCubes[color] {
				maxCubes[color] = count
			}
		}
	}

	return maxCubes["red"] <= 12 && maxCubes["green"] <= 13 && maxCubes["blue"] <= 14
}

func parseCubes(subset string) map[string]int {
	cubes := map[string]int{}
	details := strings.Split(subset, ", ")
	for _, detail := range details {
		parts := strings.Fields(detail)
		count, _ := strconv.Atoi(parts[0])
		color := parts[1]
		cubes[color] += count
	}
	return cubes
}

func getGameID(line string) int {
	parts := strings.Split(line, ":")
	id, _ := strconv.Atoi(strings.TrimSpace(strings.Split(parts[0], " ")[1]))
	return id
}

func sum(slice []int) int {
	total := 0
	for _, value := range slice {
		total += value
	}
	return total
}

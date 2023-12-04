package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Scratchcard struct {
	winningNumbers map[int]bool
	yourNumbers    []int
}

// Only has part2
func main() {
	file, err := os.Open("day4_test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var scratchcards []Scratchcard
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			continue
		}

		winningNumbersMap := make(map[int]bool)
		for _, numStr := range strings.Fields(strings.Split(parts[1], "|")[0]) {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				continue
			}
			winningNumbersMap[num] = true
		}

		yourNumbersSlice := []int{}
		for _, numStr := range strings.Fields(strings.Split(parts[1], "|")[1]) {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				continue
			}
			yourNumbersSlice = append(yourNumbersSlice, num)
		}

		scratchcards = append(scratchcards, Scratchcard{winningNumbersMap, yourNumbersSlice})
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	totalScratchcards := part2(scratchcards)
	fmt.Println(totalScratchcards)
}

func part2(scratchcards []Scratchcard) int {
	cardCounters := make([]int, len(scratchcards))
	for i := range cardCounters {
		cardCounters[i] = 1
	}

	totalCards := 0
	for i := 0; i < len(cardCounters); i++ {
		for cardCounters[i] > 0 {
			matchCount := countMatches(scratchcards[i])
			totalCards++
			cardCounters[i]--

			for j := 1; j <= matchCount; j++ {
				if i+j < len(cardCounters) {
					cardCounters[i+j]++
				}
			}
		}
	}

	return totalCards
}

func countMatches(card Scratchcard) int {
	matchCount := 0
	for _, num := range card.yourNumbers {
		if card.winningNumbers[num] {
			matchCount++
		}
	}
	return matchCount
}

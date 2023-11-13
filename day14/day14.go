package day14

import (
	"fmt"
	"github.com/the-medo/go-advent-2021/utils"
	"strings"
)

type Counter map[rune]int
type InsertionRules map[rune]map[rune]rune

func Solve(input string) {
	sections := utils.SplitByEmptyRow(input)
	mapOfInsertions := make(map[rune]map[rune]rune)

	template, insertionRuleStrings, stepCount := []rune(sections[0]), utils.SplitRows(sections[1]), 20

	for _, insertionRule := range insertionRuleStrings {
		split := strings.Split(insertionRule, " -> ")

		rune1 := rune(split[0][0])
		rune2 := rune(split[0][1])

		if mapOfInsertions[rune1] == nil {
			mapOfInsertions[rune1] = make(map[rune]rune)
		}

		mapOfInsertions[rune1][rune2] = rune(split[1][0])
	}

	fmt.Println("Template: ", template)
	fmt.Println("Map of insertions", mapOfInsertions)

	for i := 1; i <= stepCount; i++ {
		fmt.Print("Step ", i, ": ")
		template = processStep(template, mapOfInsertions)
		countRunes(template)
	}

	_, min, max := countRunes(template)
	fmt.Println("Part 1: ", max, " - ", min, " = ", max-min)

}

func processStep(runes []rune, insertionRules InsertionRules) []rune {
	newArray := make([]rune, len(runes)*2-1)

	for i, r := range runes {
		newArray[(i * 2)] = r
		if i >= len(runes)-1 {
			continue
		}
		newArray[(i*2)+1] = insertionRules[r][runes[i+1]]
	}

	return newArray
}

func countRunes(runes []rune) (Counter, int, int) {
	counter := make(Counter, 26)
	for _, r := range runes {
		counter[r]++
	}

	min, max := 0, 0
	i := 0
	for _, count := range counter {
		if i == 0 {
			min, max = count, count
		}
		if min > count && count > 0 {
			min = count
		} else if max < count {
			max = count
		}
		i++
	}

	for key, r := range counter {
		fmt.Print(string(key), "=", r, " ; ")
	}
	fmt.Println()

	return counter, min, max
}

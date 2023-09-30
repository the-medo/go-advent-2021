package main

import (
	"flag"
	"fmt"
	"github.com/the-medo/go-advent-2021/day1"
	"github.com/the-medo/go-advent-2021/utils"
	"path"
)

func main() {
	day := flag.Int("day", 1, "Which day's solution to run")
	useTestInput := flag.Bool("test", false, "Use test input")
	flag.Parse()

	fileName := fmt.Sprintf("input_%s_%d.txt", getInputType(*useTestInput), *day)
	filePath := path.Join(fmt.Sprintf("day%d", *day), fileName)

	inputData := utils.ReadFile(filePath)

	switch *day {
	case 1:
		day1.Solve(inputData)
		// ... other days
	}
}

func getInputType(isTest bool) string {
	if isTest {
		return "test"
	}
	return "real"
}

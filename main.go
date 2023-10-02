package main

import (
	"flag"
	"fmt"

	"github.com/the-medo/go-advent-2021/utils"
	"path"

	"github.com/the-medo/go-advent-2021/day1"
	"github.com/the-medo/go-advent-2021/day10"
	"github.com/the-medo/go-advent-2021/day11"
	"github.com/the-medo/go-advent-2021/day12"
	"github.com/the-medo/go-advent-2021/day13"
	"github.com/the-medo/go-advent-2021/day14"
	"github.com/the-medo/go-advent-2021/day15"
	"github.com/the-medo/go-advent-2021/day16"
	"github.com/the-medo/go-advent-2021/day17"
	"github.com/the-medo/go-advent-2021/day18"
	"github.com/the-medo/go-advent-2021/day19"
	"github.com/the-medo/go-advent-2021/day2"
	"github.com/the-medo/go-advent-2021/day20"
	"github.com/the-medo/go-advent-2021/day21"
	"github.com/the-medo/go-advent-2021/day22"
	"github.com/the-medo/go-advent-2021/day23"
	"github.com/the-medo/go-advent-2021/day24"
	"github.com/the-medo/go-advent-2021/day25"
	"github.com/the-medo/go-advent-2021/day3"
	"github.com/the-medo/go-advent-2021/day4"
	"github.com/the-medo/go-advent-2021/day5"
	"github.com/the-medo/go-advent-2021/day6"
	"github.com/the-medo/go-advent-2021/day7"
	"github.com/the-medo/go-advent-2021/day8"
	"github.com/the-medo/go-advent-2021/day9"
)

func main() {
	day := flag.Int("day", 1, "Which day's solution to run")
	useRealInput := flag.Bool("real", false, "Use real input")
	useAiSolution := flag.Bool("ai", false, "Use ai-modified solution")
	flag.Parse()

	fileName := fmt.Sprintf("input_%s_%d.txt", getInputType(*useRealInput), *day)
	filePath := path.Join(fmt.Sprintf("day%d", *day), fileName)

	inputData := utils.ReadFile(filePath)

	fmt.Printf("======= Running Day %d =========", *day)
	fmt.Println()

	switch *day {
	case 1:
		day1.Solve(inputData)
	case 2:
		day2.Solve(inputData)
	case 3:
		if *useAiSolution {
			day3.SolveAi(inputData)
			break
		} else {
			day3.Solve(inputData)
		}
	case 4:
		day4.Solve(inputData)
	case 5:
		day5.Solve(inputData)
	case 6:
		day6.Solve(inputData)
	case 7:
		day7.Solve(inputData)
	case 8:
		day8.Solve(inputData)
	case 9:
		day9.Solve(inputData)
	case 10:
		day10.Solve(inputData)
	case 11:
		day11.Solve(inputData)
	case 12:
		day12.Solve(inputData)
	case 13:
		day13.Solve(inputData)
	case 14:
		day14.Solve(inputData)
	case 15:
		day15.Solve(inputData)
	case 16:
		day16.Solve(inputData)
	case 17:
		day17.Solve(inputData)
	case 18:
		day18.Solve(inputData)
	case 19:
		day19.Solve(inputData)
	case 20:
		day20.Solve(inputData)
	case 21:
		day21.Solve(inputData)
	case 22:
		day22.Solve(inputData)
	case 23:
		day23.Solve(inputData)
	case 24:
		day24.Solve(inputData)
	case 25:
		day25.Solve(inputData)
	}
}

func getInputType(isReal bool) string {
	if isReal {
		return "real"
	}
	return "test"
}

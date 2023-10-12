package day12

import (
	"fmt"
	"github.com/the-medo/go-advent-2021/utils"
	"strings"
	"unicode"
)

type CaveRoom struct {
	Name       string
	Neighbors  []*CaveRoom
	IsSmall    bool
	IsEnd      bool
	IsUnusable bool
}

type CaveGraph = map[string]*CaveRoom

func Solve(input string) {
	rows := utils.SplitRows(input)
	caveRooms := make(CaveGraph)

	startRoom := &CaveRoom{
		Name:    "start",
		IsSmall: true,
	}

	endRoom := &CaveRoom{
		Name:    "end",
		IsSmall: true,
		IsEnd:   true,
	}

	caveRooms["start"] = startRoom
	caveRooms["end"] = endRoom

	for _, row := range rows {
		splitName := strings.Split(row, "-")
		room1, room2 := splitName[0], splitName[1]
		checkRoom(room1, caveRooms)
		checkRoom(room2, caveRooms)
		caveRooms[room1].Neighbors = append(caveRooms[room1].Neighbors, caveRooms[room2])
		caveRooms[room2].Neighbors = append(caveRooms[room2].Neighbors, caveRooms[room1])
	}

	fmt.Println(" ============= Setting IsUnusable ==========")
	for _, room := range caveRooms {
		hasBigNeighbor := false
		if room.IsEnd {
			continue
		}
		for _, neighbor := range room.Neighbors {
			if !neighbor.IsSmall || neighbor.IsEnd {
				hasBigNeighbor = true
				break
			}
		}
		if !hasBigNeighbor {
			fmt.Println("setting IsUnusable", room.Name)
			caveRooms[room.Name].IsUnusable = true
		}
	}
	fmt.Println(" ============= Setting IsUnusable finished ==========")

	allRoutes := traverseCave([]string{"start"}, &caveRooms)

	fmt.Println("===================================")
	for _, route := range allRoutes {
		fmt.Println(strings.Join(route, "-"))
	}

	fmt.Println("===================================")
	fmt.Println("Part 1 - found", len(allRoutes), "routes")
}

func checkRoom(s string, caveRooms CaveGraph) {
	if caveRooms[s] == nil {
		caveRooms[s] = &CaveRoom{Name: s}
		if unicode.IsLower(rune(s[0])) {
			caveRooms[s].IsSmall = true
		}
	}
}

func traverseCave(visited []string, graph *CaveGraph) [][]string {
	node := (*graph)[visited[len(visited)-1]]
	if node.IsEnd {
		fmt.Println("found end", visited)
		return [][]string{visited}
	}
	allRoutes := make([][]string, 0)
	for _, neighbor := range node.Neighbors {
		alreadyVisited := utils.SliceContians(visited, neighbor.Name)
		//fmt.Println("visited", visited, "alreadyVisited", alreadyVisited, neighbor.Name)
		if (neighbor.IsSmall && !alreadyVisited && !neighbor.IsUnusable) || !neighbor.IsSmall {
			newVisited := append(visited, neighbor.Name)
			allRoutes = append(allRoutes, traverseCave(newVisited, graph)...)
		}
	}

	return allRoutes
}

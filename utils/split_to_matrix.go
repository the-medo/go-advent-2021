package utils

import "strings"

func SplitToMatrix(s string) ([][]int, int) {
	boardSum := 0
	rows := SplitRows(s)

	rsp := make([][]int, len(rows))

	for i, row := range rows {
		rsp[i] = StringsToInts(strings.Fields(row))
		for _, cell := range rsp[i] {
			boardSum += cell
		}
	}

	return rsp, boardSum
}

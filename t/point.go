package t

import (
	"github.com/the-medo/go-advent-2021/utils"
	"strconv"
	"strings"
)

type Point2D struct {
	X int
	Y int
}

func (p Point2D) ToString() string {
	return strconv.Itoa(p.X) + ";" + strconv.Itoa(p.Y)
}

func (p Point2D) SurroundingPoints(includeDiagonal bool, matrixWidth int, matrixHeight int) []Point2D {
	surroundingPoints := make([]Point2D, 0)

	if p.X > 0 {
		surroundingPoints = append(surroundingPoints, Point2D{X: p.X - 1, Y: p.Y})
		if p.Y > 0 && includeDiagonal {
			surroundingPoints = append(surroundingPoints, Point2D{X: p.X - 1, Y: p.Y - 1})
		}
		if p.Y < matrixHeight-1 && includeDiagonal {
			surroundingPoints = append(surroundingPoints, Point2D{X: p.X - 1, Y: p.Y + 1})
		}
	}
	if p.X < matrixWidth-1 {
		surroundingPoints = append(surroundingPoints, Point2D{X: p.X + 1, Y: p.Y})
		if p.Y > 0 && includeDiagonal {
			surroundingPoints = append(surroundingPoints, Point2D{X: p.X + 1, Y: p.Y - 1})
		}
		if p.Y < matrixHeight-1 && includeDiagonal {
			surroundingPoints = append(surroundingPoints, Point2D{X: p.X + 1, Y: p.Y + 1})
		}
	}
	if p.Y > 0 {
		surroundingPoints = append(surroundingPoints, Point2D{X: p.X, Y: p.Y - 1})
	}
	if p.Y < matrixHeight-1 {
		surroundingPoints = append(surroundingPoints, Point2D{X: p.X, Y: p.Y + 1})
	}
	return surroundingPoints
}

func StringToPoint2D(s string) Point2D {
	coords := utils.StringsToInts(strings.Split(s, ";"))
	return Point2D{X: coords[0], Y: coords[1]}
}

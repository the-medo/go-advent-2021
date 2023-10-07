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

func StringToPoint2D(s string) Point2D {
	coords := utils.StringsToInts(strings.Split(s, ";"))
	return Point2D{X: coords[0], Y: coords[1]}
}
